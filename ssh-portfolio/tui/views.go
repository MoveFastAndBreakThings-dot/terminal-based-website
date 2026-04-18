package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"ssh-portfolio/content"
)

// ── Fixed chrome ─────────────────────────────────────────────────────────────

func renderHeader(m Model, s Styles) string {
	left := s.Header.Render("  v1.0.0")

	bracketAccent := lipgloss.NewStyle().Foreground(m.theme.Accent)
	bracketMuted := lipgloss.NewStyle().Foreground(m.theme.Muted)
	right := bracketAccent.Render("[") +
		bracketMuted.Render("   ssh -p 23234 samar-personal-portfolio.fly.dev  ") +
		bracketAccent.Render("]")

	lw := lipgloss.Width(left)
	rw := lipgloss.Width(right)
	gap := m.width - lw - rw
	if gap < 0 {
		gap = 0
	}
	return left + strings.Repeat(" ", gap) + right
}

func renderTabBar(m Model, s Styles) string {
	var parts []string
	for i, label := range m.tabs {
		if i == m.activeTab {
			parts = append(parts, s.TabActive.Render(label))
		} else {
			parts = append(parts, s.TabInactive.Render(label))
		}
	}
	tabLine := "  " + strings.Join(parts, "   ")
	rule := lipgloss.NewStyle().Foreground(m.theme.Dim).Render(
		strings.Repeat("─", m.width),
	)
	return tabLine + "\n" + rule
}

func renderFooter(m Model, s Styles) string {
	return s.Footer.Render("  ↑↓ scroll   ←→ / hl tabs   1-5 jump   t theme   q quit")
}

// ── Tab dispatcher ────────────────────────────────────────────────────────────

func renderActiveTab(m Model) string {
	switch m.activeTab {
	case 0:
		return renderAbout(m)
	case 1:
		return renderExperience(m)
	case 2:
		return renderProjects(m)
	case 3:
		return renderSkills(m)
	case 4:
		return renderLinks(m)
	}
	return ""
}

// ── About ─────────────────────────────────────────────────────────────────────

func renderAbout(m Model) string {
	nameStyle := lipgloss.NewStyle().Foreground(m.theme.Accent).Bold(true)
	roleStyle := lipgloss.NewStyle().Foreground(m.theme.Muted)
	dimStyle := lipgloss.NewStyle().Foreground(m.theme.Dim)
	hiStyle := lipgloss.NewStyle().Foreground(m.theme.Accent).Bold(true)

	imgBlock := GetAscii(m.theme)
	imgW := lipgloss.Width(imgBlock)

	const gap = 3
	textW := m.width - imgW - gap - 2
	if textW >= 30 {
		bioStyle := lipgloss.NewStyle().Foreground(m.theme.Foreground).Width(textW)
		var textLines []string
		textLines = append(textLines, nameStyle.Render(content.MyProfile.Name))
		textLines = append(textLines, roleStyle.Render(content.MyProfile.Role))
		textLines = append(textLines, "")
		for i, para := range content.MyProfile.Bio {
			textLines = append(textLines, bioStyle.Render(para))
			if i < len(content.MyProfile.Bio)-1 {
				textLines = append(textLines, "")
			}
		}
		textLines = append(textLines, "")
		textLines = append(textLines, hiStyle.Render("Highlights"))
		for _, ec := range content.Extracurriculars {
			textLines = append(textLines, dimStyle.Render("· ")+dimStyle.Render(ec))
		}
		return lipgloss.JoinHorizontal(lipgloss.Top, imgBlock, strings.Repeat(" ", gap), strings.Join(textLines, "\n"))
	}

	// narrow terminal: stack
	bioWidth := m.width - 4
	if bioWidth < 30 {
		bioWidth = 30
	}
	bioStyle := lipgloss.NewStyle().Foreground(m.theme.Foreground).Width(bioWidth)
	var lines []string
	lines = append(lines, imgBlock)
	lines = append(lines, "")
	lines = append(lines, "  "+nameStyle.Render(content.MyProfile.Name))
	lines = append(lines, "  "+roleStyle.Render(content.MyProfile.Role))
	lines = append(lines, "")
	for i, para := range content.MyProfile.Bio {
		lines = append(lines, "  "+bioStyle.Render(para))
		if i < len(content.MyProfile.Bio)-1 {
			lines = append(lines, "")
		}
	}
	lines = append(lines, "")
	lines = append(lines, "  "+hiStyle.Render("Highlights"))
	for _, ec := range content.Extracurriculars {
		lines = append(lines, "  "+dimStyle.Render("·")+" "+dimStyle.Render(ec))
	}
	return strings.Join(lines, "\n")
}

// ── Experience ────────────────────────────────────────────────────────────────

func renderExperience(m Model) string {
	s := NewStyles(m.theme)
	bulletAccent := lipgloss.NewStyle().Foreground(m.theme.Accent).Render("›")
	dimRule := lipgloss.NewStyle().Foreground(m.theme.Dim).Render(
		strings.Repeat("─", m.width-4),
	)

	bulletWidth := m.width - 8
	if bulletWidth < 20 {
		bulletWidth = 20
	}
	bulletStyle := lipgloss.NewStyle().Foreground(m.theme.Foreground).Width(bulletWidth)

	var sections []string
	for _, job := range content.Jobs {
		var lines []string

		// Title line: job title left, period right-aligned
		titleStr := s.JobTitle.Render(job.Title)
		periodStr := s.DateText.Render(job.Period)
		gap := m.width - 4 - lipgloss.Width(titleStr) - lipgloss.Width(periodStr)
		if gap < 1 {
			gap = 1
		}
		lines = append(lines, "  "+titleStr+strings.Repeat(" ", gap)+periodStr)
		lines = append(lines, "  "+s.CompanyText.Render(job.Company))
		lines = append(lines, "")

		for _, b := range job.Bullets {
			lines = append(lines, "  "+bulletAccent+" "+bulletStyle.Render(b))
		}
		lines = append(lines, "")

		// Tech bullets
		tagStyle := lipgloss.NewStyle().Foreground(m.theme.Muted)
		tagAccent := lipgloss.NewStyle().Foreground(m.theme.Accent)
		for _, t := range job.Tags {
			lines = append(lines, "    "+tagAccent.Render("•")+" "+tagStyle.Render(t))
		}

		sections = append(sections, strings.Join(lines, "\n"))
	}

	return strings.Join(sections, "\n\n  "+dimRule+"\n\n")
}

// ── Projects ──────────────────────────────────────────────────────────────────

func renderProjects(m Model) string {
	nameStyle := lipgloss.NewStyle().Foreground(m.theme.Accent).Bold(true)
	eventStyle := lipgloss.NewStyle().Foreground(m.theme.Dim)
	arrowStyle := lipgloss.NewStyle().Foreground(m.theme.Accent)
	linkStyle := lipgloss.NewStyle().Foreground(m.theme.Link)

	descWidth := m.width - 4
	if descWidth < 20 {
		descWidth = 20
	}
	descStyle := lipgloss.NewStyle().Foreground(m.theme.Foreground).Width(descWidth)

	var sections []string
	for _, proj := range content.Projects {
		var lines []string

		header := nameStyle.Render(proj.Name)
		if proj.Event != "" {
			header += "  " + eventStyle.Render("["+proj.Event+"]")
		}
		lines = append(lines, "  "+header)
		lines = append(lines, "  "+descStyle.Render(proj.Description))

		tagStyle := lipgloss.NewStyle().Foreground(m.theme.Muted)
		tagAccentStyle := lipgloss.NewStyle().Foreground(m.theme.Accent)
		for _, t := range proj.Tags {
			lines = append(lines, "    "+tagAccentStyle.Render("•")+" "+tagStyle.Render(t))
		}
		if proj.URL != "" {
			arrow := arrowStyle.Render("→")
			urlRendered := linkStyle.Render(osc8Link(proj.URL, proj.URL))
			lines = append(lines, "  "+arrow+"  "+urlRendered)
		}

		sections = append(sections, strings.Join(lines, "\n"))
	}
	return strings.Join(sections, "\n\n")
}

// ── Skills ────────────────────────────────────────────────────────────────────

func renderSkills(m Model) string {
	_ = NewStyles(m.theme)
	catStyle := lipgloss.NewStyle().Foreground(m.theme.Accent).Bold(true)
	itemStyle := lipgloss.NewStyle().Foreground(m.theme.Foreground)
	dotStyle := lipgloss.NewStyle().Foreground(m.theme.Accent)

	var lines []string
	for _, group := range content.SkillGroups {
		lines = append(lines, "  "+catStyle.Render(group.Category))
		for _, item := range group.Items {
			lines = append(lines, "    "+dotStyle.Render("•")+" "+itemStyle.Render(item))
		}
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n")
}

// ── Links ─────────────────────────────────────────────────────────────────────

func renderLinks(m Model) string {
	s := NewStyles(m.theme)
	arrowStyle := lipgloss.NewStyle().Foreground(m.theme.Accent)
	labelStyle := lipgloss.NewStyle().Foreground(m.theme.Muted).Width(12)
	noteStyle := lipgloss.NewStyle().Foreground(m.theme.Dim)

	var lines []string
	lines = append(lines, "")

	for _, lk := range content.Links {
		arrow := arrowStyle.Render("→")
		label := labelStyle.Render(lk.Label)

		displayURL := lk.URL
		if strings.HasPrefix(lk.URL, "mailto:") {
			displayURL = strings.TrimPrefix(lk.URL, "mailto:")
		}

		// OSC 8 hyperlink wrapping the lipgloss-styled display text
		styledURL := s.LinkText.Render(displayURL)
		hyperlink := osc8Link(lk.URL, styledURL)

		lines = append(lines, fmt.Sprintf("  %s  %s  %s", arrow, label, hyperlink))
	}

	lines = append(lines, "")
	lines = append(lines, "  "+noteStyle.Render("( links are clickable in iTerm2, WezTerm, Kitty )"))
	return strings.Join(lines, "\n")
}

// ── Helpers ───────────────────────────────────────────────────────────────────

// osc8Link wraps displayText in an OSC 8 hyperlink escape for terminals that
// support it (iTerm2, WezTerm, Kitty, …).
func osc8Link(url, displayText string) string {
	return "\033]8;;" + url + "\033\\" + displayText + "\033]8;;\033\\"
}
