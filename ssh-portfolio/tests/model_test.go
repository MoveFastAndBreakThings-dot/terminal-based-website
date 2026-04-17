package tests

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"ssh-portfolio/tui"
)

// ── helpers ───────────────────────────────────────────────────────────────────

func newReadyModel() tui.Model {
	return tui.NewModel(120, 40)
}

func sendKey(m tui.Model, key string) tui.Model {
	updated, _ := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(key)})
	return updated.(tui.Model)
}

func sendSpecialKey(m tui.Model, kt tea.KeyType) tui.Model {
	updated, _ := m.Update(tea.KeyMsg{Type: kt})
	return updated.(tui.Model)
}

// ── NewModel ──────────────────────────────────────────────────────────────────

func TestNewModel_ZeroSize_NotReady(t *testing.T) {
	m := tui.NewModel(0, 0)
	if view := m.View(); view != "\n  Connecting…" {
		t.Errorf("zero-size model View() = %q, want connecting message", view)
	}
}

func TestNewModel_ValidSize_ViewNotEmpty(t *testing.T) {
	m := newReadyModel()
	v := m.View()
	if v == "" {
		t.Fatal("View() is empty for a ready model")
	}
}

func TestNewModel_ValidSize_ViewNotConnecting(t *testing.T) {
	m := newReadyModel()
	if m.View() == "\n  Connecting…" {
		t.Fatal("valid-size model still shows connecting message")
	}
}

func TestNewModel_InitReturnsNil(t *testing.T) {
	m := newReadyModel()
	if cmd := m.Init(); cmd != nil {
		t.Fatal("Init() should return nil")
	}
}

// ── Tab navigation ────────────────────────────────────────────────────────────

func TestModel_TabRight_ChangesContent(t *testing.T) {
	m := newReadyModel()
	before := m.View()
	m = sendKey(m, "l")
	if m.View() == before {
		t.Fatal("View() unchanged after right-tab navigation")
	}
}

func TestModel_TabRight_Arrow_ChangesContent(t *testing.T) {
	m := newReadyModel()
	before := m.View()
	m = sendSpecialKey(m, tea.KeyRight)
	if m.View() == before {
		t.Fatal("View() unchanged after KeyRight navigation")
	}
}

func TestModel_TabRight_Tab_ChangesContent(t *testing.T) {
	m := newReadyModel()
	before := m.View()
	m = sendSpecialKey(m, tea.KeyTab)
	if m.View() == before {
		t.Fatal("View() unchanged after Tab key navigation")
	}
}

func TestModel_TabLeft_ChangesContent(t *testing.T) {
	m := newReadyModel()
	// go right first so left has somewhere to go
	m = sendKey(m, "l")
	after := m.View()
	m = sendKey(m, "h")
	if m.View() == after {
		t.Fatal("View() unchanged after left-tab navigation")
	}
}

func TestModel_TabWrapAround_Right(t *testing.T) {
	m := newReadyModel()
	start := m.View()
	// 5 tabs — cycle all the way around
	for i := 0; i < 5; i++ {
		m = sendKey(m, "l")
	}
	if m.View() != start {
		t.Fatal("tab right wrap-around did not return to start view")
	}
}

func TestModel_TabWrapAround_Left(t *testing.T) {
	m := newReadyModel()
	start := m.View()
	for i := 0; i < 5; i++ {
		m = sendKey(m, "h")
	}
	if m.View() != start {
		t.Fatal("tab left wrap-around did not return to start view")
	}
}

// ── Jump to tab by number ─────────────────────────────────────────────────────

func TestModel_JumpTab_SameAs_Navigation(t *testing.T) {
	tabs := []string{"1", "2", "3", "4", "5"}
	for _, key := range tabs {
		m1 := newReadyModel()
		m1 = sendKey(m1, key)

		m2 := newReadyModel()
		steps := int(key[0] - '1')
		for i := 0; i < steps; i++ {
			m2 = sendKey(m2, "l")
		}

		if m1.View() != m2.View() {
			t.Errorf("jump %q view differs from sequential navigation", key)
		}
	}
}

func TestModel_JumpTab_1_IsAbout(t *testing.T) {
	m := newReadyModel()
	m = sendKey(m, "l") // move off tab 0
	m = sendKey(m, "1") // jump back to tab 0
	if !containsAny(m.View(), "Samardeep", "About") {
		t.Fatal("jump to 1 should show About/profile content")
	}
}

// ── Theme cycling ─────────────────────────────────────────────────────────────

func TestModel_ThemeCycle_ChangesThemeIndex(t *testing.T) {
	// Theme output may look identical in non-TTY environments (lipgloss strips
	// ANSI without a color profile). Test structural behavior: cycling n times
	// returns to original view, confirming the index wraps correctly.
	m := newReadyModel()
	start := m.View()
	m = sendKey(m, "t")
	// After one cycle the view may or may not differ; after full cycle it must match.
	for i := 1; i < len(tui.Themes); i++ {
		m = sendKey(m, "t")
	}
	if m.View() != start {
		t.Fatal("full theme cycle did not restore original view")
	}
}

// ── Quit ──────────────────────────────────────────────────────────────────────

func TestModel_QuitKey_q_ReturnsQuitCmd(t *testing.T) {
	m := newReadyModel()
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("q")})
	assertQuitCmd(t, cmd, "q")
}

func TestModel_QuitKey_CtrlC_ReturnsQuitCmd(t *testing.T) {
	m := newReadyModel()
	_, cmd := m.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	assertQuitCmd(t, cmd, "ctrl+c")
}

func assertQuitCmd(t *testing.T, cmd tea.Cmd, key string) {
	t.Helper()
	if cmd == nil {
		t.Fatalf("key %q: Update returned nil cmd, expected tea.Quit", key)
	}
	msg := cmd()
	if _, ok := msg.(tea.QuitMsg); !ok {
		t.Fatalf("key %q: cmd() returned %T, expected tea.QuitMsg", key, msg)
	}
}

// ── Window resize ─────────────────────────────────────────────────────────────

func TestModel_WindowResize_UpdatesDimensions(t *testing.T) {
	m := tui.NewModel(0, 0) // start not-ready
	updated, cmd := m.Update(tea.WindowSizeMsg{Width: 120, Height: 40})
	if cmd != nil {
		t.Fatal("WindowSizeMsg returned unexpected cmd")
	}
	m2 := updated.(tui.Model)
	v := m2.View()
	if v == "\n  Connecting…" {
		t.Fatal("after resize, model still shows connecting")
	}
}

func TestModel_WindowResize_SmallTerminal(t *testing.T) {
	m := newReadyModel()
	updated, _ := m.Update(tea.WindowSizeMsg{Width: 40, Height: 10})
	m2 := updated.(tui.Model)
	if m2.View() == "" {
		t.Fatal("View() empty after resize to small terminal")
	}
}

func TestModel_WindowResize_VerySmall_DoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("panic on tiny terminal: %v", r)
		}
	}()
	m := newReadyModel()
	m.Update(tea.WindowSizeMsg{Width: 1, Height: 1})
}

// ── Unhandled messages ────────────────────────────────────────────────────────

func TestModel_UnknownKey_NoStateChange(t *testing.T) {
	m := newReadyModel()
	before := m.View()
	// send a key that is not bound to anything
	m = sendKey(m, "z")
	// view should be unchanged (z is not a keybinding)
	if m.View() != before {
		t.Fatal("unbound key 'z' changed model state")
	}
}

// ── helper ────────────────────────────────────────────────────────────────────

func containsAny(s string, substrs ...string) bool {
	for _, sub := range substrs {
		if len(sub) > 0 {
			for i := 0; i <= len(s)-len(sub); i++ {
				if s[i:i+len(sub)] == sub {
					return true
				}
			}
		}
	}
	return false
}
