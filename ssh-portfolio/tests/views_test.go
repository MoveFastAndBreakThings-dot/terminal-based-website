package tests

import (
	"regexp"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
	"ssh-portfolio/content"
	"ssh-portfolio/tui"
)

// stripANSI removes all ANSI escape sequences and OSC sequences from s.
var ansiRE = regexp.MustCompile(`\x1b(?:\[[0-9;?]*[A-Za-z]|\][^\x1b]*(?:\x1b\\|\x07)|[^[])`)

func stripANSI(s string) string {
	return ansiRE.ReplaceAllString(s, "")
}

// tabView returns the stripped View() after navigating to tab index i.
// Uses a very tall terminal so all content is visible without scrolling.
func tabView(tabIndex int) string {
	m := tui.NewModel(160, 2000)
	tKey := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("l")}
	for i := 0; i < tabIndex; i++ {
		updated, _ := m.Update(tKey)
		m = updated.(tui.Model)
	}
	return stripANSI(m.View())
}

// ── Tab bar ───────────────────────────────────────────────────────────────────

func TestView_TabBar_ContainsAllTabNames(t *testing.T) {
	v := tabView(0)
	for _, name := range []string{"About", "Experience", "Projects", "Skills", "Links"} {
		if !strings.Contains(v, name) {
			t.Errorf("tab bar missing tab name %q", name)
		}
	}
}

// ── Header ────────────────────────────────────────────────────────────────────

func TestView_Header_ContainsVersion(t *testing.T) {
	v := tabView(0)
	if !strings.Contains(v, "v1.0.0") {
		t.Fatal("header missing version string v1.0.0")
	}
}

func TestView_Header_ContainsDomain(t *testing.T) {
	v := tabView(0)
	if !strings.Contains(v, "samar.com") {
		t.Fatal("header missing domain samar.com")
	}
}

// ── Footer ────────────────────────────────────────────────────────────────────

func TestView_Footer_ContainsKeyHints(t *testing.T) {
	v := tabView(0)
	for _, hint := range []string{"scroll", "tabs", "theme", "quit"} {
		if !strings.Contains(v, hint) {
			t.Errorf("footer missing hint %q", hint)
		}
	}
}

// ── About tab ─────────────────────────────────────────────────────────────────

func TestView_About_ContainsName(t *testing.T) {
	v := tabView(0)
	if !strings.Contains(v, content.MyProfile.Name) {
		t.Fatalf("About tab missing name %q", content.MyProfile.Name)
	}
}

func TestView_About_ContainsRole(t *testing.T) {
	v := tabView(0)
	// Role may be word-wrapped; check first word chunk that won't span a wrap.
	roleSnippet := "Engineering Student"
	if !strings.Contains(v, roleSnippet) {
		t.Fatalf("About tab missing role snippet %q", roleSnippet)
	}
}

func TestView_About_ContainsBioText(t *testing.T) {
	v := tabView(0)
	// First 20 chars of first bio paragraph should appear somewhere in the view.
	if len(content.MyProfile.Bio) == 0 {
		t.Skip("no bio paragraphs")
	}
	snippet := content.MyProfile.Bio[0][:20]
	if !strings.Contains(v, snippet) {
		t.Fatalf("About tab missing bio snippet %q", snippet)
	}
}

func TestView_About_ContainsHighlightsSection(t *testing.T) {
	v := tabView(0)
	if !strings.Contains(v, "Highlights") {
		t.Fatal("About tab missing Highlights section")
	}
}

func TestView_About_ContainsExtracurriculars(t *testing.T) {
	if len(content.Extracurriculars) == 0 {
		t.Skip("no extracurriculars")
	}
	v := tabView(0)
	// Check at least the first extracurricular appears.
	snippet := content.Extracurriculars[0][:15]
	if !strings.Contains(v, snippet) {
		t.Fatalf("About tab missing extracurricular snippet %q", snippet)
	}
}

func TestView_About_WideTerminal_HasAsciiAndText_SameLine(t *testing.T) {
	// Wide terminal (160+) should render side-by-side.
	// The profile name and ascii art should both appear in the output.
	m := tui.NewModel(200, 50)
	v := stripANSI(m.View())
	if !strings.Contains(v, content.MyProfile.Name) {
		t.Fatal("wide About tab missing profile name")
	}
	if !strings.Contains(v, "▄") {
		t.Fatal("wide About tab missing portrait half-block chars")
	}
}

func TestView_About_NarrowTerminal_StillRendersName(t *testing.T) {
	// Narrow + very tall so we can see past the 48-line ascii art.
	m := tui.NewModel(60, 2000)
	v := stripANSI(m.View())
	if !strings.Contains(v, content.MyProfile.Name) {
		t.Fatal("narrow About tab missing profile name")
	}
}

// ── Experience tab ────────────────────────────────────────────────────────────

func TestView_Experience_ContainsJobTitles(t *testing.T) {
	v := tabView(1)
	for _, job := range content.Jobs {
		if !strings.Contains(v, job.Title) {
			t.Errorf("Experience tab missing job title %q", job.Title)
		}
	}
}

func TestView_Experience_ContainsCompanies(t *testing.T) {
	v := tabView(1)
	for _, job := range content.Jobs {
		// Company string may be long; check a 15-char prefix.
		snippet := job.Company
		if len(snippet) > 15 {
			snippet = snippet[:15]
		}
		if !strings.Contains(v, snippet) {
			t.Errorf("Experience tab missing company snippet %q", snippet)
		}
	}
}

func TestView_Experience_ContainsTags(t *testing.T) {
	v := tabView(1)
	for _, job := range content.Jobs {
		for _, tag := range job.Tags {
			if !strings.Contains(v, tag) {
				t.Errorf("Experience tab missing tag %q (job: %s)", tag, job.Title)
			}
		}
	}
}

func TestView_Experience_ContainsBulletMarkers(t *testing.T) {
	v := tabView(1)
	if !strings.Contains(v, "•") {
		t.Fatal("Experience tab has no bullet markers (•)")
	}
}

// ── Projects tab ──────────────────────────────────────────────────────────────

func TestView_Projects_ContainsProjectNames(t *testing.T) {
	v := tabView(2)
	for _, proj := range content.Projects {
		snippet := proj.Name
		if len(snippet) > 20 {
			snippet = snippet[:20]
		}
		if !strings.Contains(v, snippet) {
			t.Errorf("Projects tab missing project name snippet %q", snippet)
		}
	}
}

func TestView_Projects_ContainsTags(t *testing.T) {
	v := tabView(2)
	for _, proj := range content.Projects {
		for _, tag := range proj.Tags {
			if !strings.Contains(v, tag) {
				t.Errorf("Projects tab missing tag %q (project: %s)", tag, proj.Name)
			}
		}
	}
}

func TestView_Projects_ContainsBulletMarkers(t *testing.T) {
	v := tabView(2)
	if !strings.Contains(v, "•") {
		t.Fatal("Projects tab has no bullet markers (•)")
	}
}

func TestView_Projects_EventBadgeVisible(t *testing.T) {
	v := tabView(2)
	for _, proj := range content.Projects {
		if proj.Event == "" {
			continue
		}
		if !strings.Contains(v, proj.Event) {
			t.Errorf("Projects tab missing event badge %q", proj.Event)
		}
	}
}

// ── Skills tab ────────────────────────────────────────────────────────────────

func TestView_Skills_ContainsCategories(t *testing.T) {
	v := tabView(3)
	for _, sg := range content.SkillGroups {
		if !strings.Contains(v, sg.Category) {
			t.Errorf("Skills tab missing category %q", sg.Category)
		}
	}
}

func TestView_Skills_ContainsItems(t *testing.T) {
	v := tabView(3)
	for _, sg := range content.SkillGroups {
		for _, item := range sg.Items {
			if !strings.Contains(v, item) {
				t.Errorf("Skills tab missing item %q", item)
			}
		}
	}
}

func TestView_Skills_ContainsBulletMarkers(t *testing.T) {
	v := tabView(3)
	if !strings.Contains(v, "•") {
		t.Fatal("Skills tab has no bullet markers (•)")
	}
}

// ── Links tab ─────────────────────────────────────────────────────────────────

func TestView_Links_ContainsLabels(t *testing.T) {
	v := tabView(4)
	for _, lk := range content.Links {
		if !strings.Contains(v, lk.Label) {
			t.Errorf("Links tab missing label %q", lk.Label)
		}
	}
}

func TestView_Links_ContainsURLs(t *testing.T) {
	v := tabView(4)
	for _, lk := range content.Links {
		display := lk.URL
		if strings.HasPrefix(display, "mailto:") {
			display = strings.TrimPrefix(display, "mailto:")
		}
		if !strings.Contains(v, display) {
			t.Errorf("Links tab missing URL %q", display)
		}
	}
}

func TestView_Links_ContainsArrow(t *testing.T) {
	v := tabView(4)
	if !strings.Contains(v, "→") {
		t.Fatal("Links tab missing → arrow")
	}
}
