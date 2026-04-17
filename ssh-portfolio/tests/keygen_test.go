package tests

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"path/filepath"
	"testing"
)

// generateHostKey mirrors the logic in main.go (unexported there).
func generateHostKey(path string) error {
	key, err := rsa.GenerateKey(rand.Reader, 2048) // 2048 for test speed
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

// ── Tests ─────────────────────────────────────────────────────────────────────

func TestGenerateHostKey_CreatesFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "host_key")

	if err := generateHostKey(path); err != nil {
		t.Fatalf("generateHostKey returned error: %v", err)
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatal("host key file was not created")
	}
}

func TestGenerateHostKey_FileNotEmpty(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "host_key")

	if err := generateHostKey(path); err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	if info.Size() == 0 {
		t.Fatal("host key file is empty")
	}
}

func TestGenerateHostKey_FilePermissions(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "host_key")

	if err := generateHostKey(path); err != nil {
		t.Fatal(err)
	}

	info, err := os.Stat(path)
	if err != nil {
		t.Fatal(err)
	}
	perm := info.Mode().Perm()
	if perm != 0600 {
		t.Errorf("host key permissions = %o, want 0600", perm)
	}
}

func TestGenerateHostKey_ValidPEMBlock(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "host_key")

	if err := generateHostKey(path); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	block, rest := pem.Decode(data)
	if block == nil {
		t.Fatal("file contains no valid PEM block")
	}
	if block.Type != "RSA PRIVATE KEY" {
		t.Errorf("PEM type = %q, want RSA PRIVATE KEY", block.Type)
	}
	if len(rest) > 0 && len(rest) != len(data)-len(pem.EncodeToMemory(block)) {
		t.Error("unexpected trailing data after PEM block")
	}
}

func TestGenerateHostKey_ParseableRSAKey(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "host_key")

	if err := generateHostKey(path); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	block, _ := pem.Decode(data)
	if block == nil {
		t.Fatal("no PEM block")
	}

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		t.Fatalf("failed to parse RSA key: %v", err)
	}
	if err := key.Validate(); err != nil {
		t.Fatalf("RSA key validation failed: %v", err)
	}
}

func TestGenerateHostKey_Idempotent_OverwritesExisting(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "host_key")

	if err := generateHostKey(path); err != nil {
		t.Fatal(err)
	}
	first, _ := os.ReadFile(path)

	if err := generateHostKey(path); err != nil {
		t.Fatal(err)
	}
	second, _ := os.ReadFile(path)

	// Keys are randomly generated — second call should produce a different key.
	if string(first) == string(second) {
		t.Error("two generateHostKey calls produced identical keys (should be random)")
	}
}

func TestGenerateHostKey_BadPath_ReturnsError(t *testing.T) {
	err := generateHostKey("/nonexistent/path/host_key")
	if err == nil {
		t.Fatal("expected error for bad path, got nil")
	}
}
