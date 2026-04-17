package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	cssh "github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	bm "github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"

	"ssh-portfolio/tui"
)

const (
	defaultPort     = 23234
	hostKeyPath     = ".ssh/host_key"
	shutdownTimeout = 30 * time.Second
)

func main() {
	port := defaultPort
	if p := os.Getenv("SSH_PORT"); p != "" {
		if n, err := strconv.Atoi(p); err == nil {
			port = n
		}
	}

	// Ensure host-key directory exists.
	if err := os.MkdirAll(".ssh", 0700); err != nil {
		log.Fatalf("mkdir .ssh: %v", err)
	}

	// Generate host key on first run; reuse on subsequent starts.
	if _, err := os.Stat(hostKeyPath); os.IsNotExist(err) {
		if err := generateHostKey(hostKeyPath); err != nil {
			log.Fatalf("generate host key: %v", err)
		}
		log.Printf("Generated new RSA host key → %s", hostKeyPath)
	}

	srv, err := wish.NewServer(
		wish.WithAddress(fmt.Sprintf(":%d", port)),
		wish.WithHostKeyPath(hostKeyPath),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)
	if err != nil {
		log.Fatalf("create SSH server: %v", err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	log.Printf("SSH portfolio listening on :%d  →  ssh -p %d localhost", port, port)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("server stopped: %v", err)
		}
	}()

	<-done
	log.Println("Shutting down…")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("shutdown error: %v", err)
	}
	log.Println("Server stopped.")
}

// teaHandler creates a fresh TUI model for every incoming SSH connection.
func teaHandler(s cssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()
	log.Printf("connection  ip=%s  time=%s", s.RemoteAddr(), time.Now().Format(time.RFC3339))
	m := tui.NewModel(pty.Window.Width, pty.Window.Height)
	return m, []tea.ProgramOption{tea.WithAltScreen()}
}

// generateHostKey writes a fresh 4096-bit RSA private key in PEM format.
func generateHostKey(path string) error {
	key, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return pem.Encode(f, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
}
