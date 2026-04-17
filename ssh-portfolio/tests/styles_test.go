package tests

import (
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"ssh-portfolio/tui"
)

// ── Themes ────────────────────────────────────────────────────────────────────

func TestThemes_MinimumCount(t *testing.T) {
	if len(tui.Themes) < 1 {
		t.Fatal("Themes slice is empty")
	}
}

func TestThemes_HaveNames(t *testing.T) {
	for i, th := range tui.Themes {
		if th.Name == "" {
			t.Errorf("Themes[%d].Name is empty", i)
		}
	}
}

func TestThemes_NoDuplicateNames(t *testing.T) {
	seen := map[string]bool{}
	for _, th := range tui.Themes {
		if seen[th.Name] {
			t.Errorf("duplicate theme name: %q", th.Name)
		}
		seen[th.Name] = true
	}
}

func TestThemes_ColorsNotEmpty(t *testing.T) {
	for _, th := range tui.Themes {
		if th.Accent == "" {
			t.Errorf("Theme %q: Accent color empty", th.Name)
		}
		if th.Foreground == "" {
			t.Errorf("Theme %q: Foreground color empty", th.Name)
		}
		if th.Background == "" {
			t.Errorf("Theme %q: Background color empty", th.Name)
		}
		if th.Muted == "" {
			t.Errorf("Theme %q: Muted color empty", th.Name)
		}
		if th.Dim == "" {
			t.Errorf("Theme %q: Dim color empty", th.Name)
		}
		if th.Link == "" {
			t.Errorf("Theme %q: Link color empty", th.Name)
		}
		if th.Border == "" {
			t.Errorf("Theme %q: Border color empty", th.Name)
		}
	}
}

// ── NewStyles ─────────────────────────────────────────────────────────────────

func TestNewStyles_DoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("NewStyles panicked: %v", r)
		}
	}()
	for _, th := range tui.Themes {
		tui.NewStyles(th)
	}
}

func TestThemeCycle_FullCycle_RestoresView(t *testing.T) {
	// Structural test: cycling through all themes returns to original state.
	// Avoids depending on ANSI output (lipgloss suppresses colors in non-TTY).
	tKey := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("t")}
	m := tui.NewModel(120, 40)
	start := m.View()
	for range tui.Themes {
		updated, _ := m.Update(tKey)
		m = updated.(tui.Model)
	}
	if m.View() != start {
		t.Fatal("cycling through all themes did not restore original view")
	}
}

// ── GetAscii ──────────────────────────────────────────────────────────────────

func TestGetAscii_NotEmpty(t *testing.T) {
	for _, th := range tui.Themes {
		art := tui.GetAscii(th)
		if art == "" {
			t.Errorf("GetAscii returned empty string for theme %q", th.Name)
		}
	}
}

func TestGetAscii_ContainsBlockChars(t *testing.T) {
	art := tui.GetAscii(tui.Themes[0])
	found := false
	for _, ch := range art {
		if ch == '▄' || ch == '▓' || ch == '▒' || ch == '░' || ch == '█' {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("GetAscii output contains no block characters — wrong file?")
	}
}

func TestGetAscii_DifferentThemesHaveDifferentColors(t *testing.T) {
	if len(tui.Themes) < 2 {
		t.Skip("need at least 2 themes")
	}
	// Test theme color values differ — the actual ANSI output may be identical
	// in non-TTY environments where lipgloss strips color codes.
	if tui.Themes[0].Dim == tui.Themes[1].Dim {
		t.Errorf("Theme %q and %q share the same Dim color", tui.Themes[0].Name, tui.Themes[1].Name)
	}
}
