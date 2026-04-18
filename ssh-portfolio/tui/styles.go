package tui

import "github.com/charmbracelet/lipgloss"

// Theme holds all color tokens for a named palette.
type Theme struct {
	Name       string
	Background lipgloss.Color
	Foreground lipgloss.Color
	Accent     lipgloss.Color
	Muted      lipgloss.Color
	Border     lipgloss.Color
	Link       lipgloss.Color
	Dim        lipgloss.Color
}

// Themes lists the three available palettes (index matches themeIndex).
var Themes = []Theme{
	{
		Name:       "Dark",
		Background: lipgloss.Color("#0d0d0d"),
		Foreground: lipgloss.Color("#e0e0e0"),
		Accent:     lipgloss.Color("#ff6b35"),
		Muted:      lipgloss.Color("#888888"),
		Border:     lipgloss.Color("#333333"),
		Link:       lipgloss.Color("#4a9eff"),
		Dim:        lipgloss.Color("#555555"),
	},
	{
		Name:       "Light",
		Background: lipgloss.Color("#f5f5f0"),
		Foreground: lipgloss.Color("#1a1a1a"),
		Accent:     lipgloss.Color("#c0392b"),
		Muted:      lipgloss.Color("#666666"),
		Border:     lipgloss.Color("#cccccc"),
		Link:       lipgloss.Color("#2563eb"),
		Dim:        lipgloss.Color("#999999"),
	},
	{
		Name:       "Hacker",
		Background: lipgloss.Color("#000000"),
		Foreground: lipgloss.Color("#00ff41"),
		Accent:     lipgloss.Color("#00ff41"),
		Muted:      lipgloss.Color("#006600"),
		Border:     lipgloss.Color("#003300"),
		Link:       lipgloss.Color("#00ff41"),
		Dim:        lipgloss.Color("#004400"),
	},
}

// Styles is a compiled set of lipgloss styles derived from a Theme.
type Styles struct {
	TabActive    lipgloss.Style
	TabInactive  lipgloss.Style
	Header       lipgloss.Style
	Footer       lipgloss.Style
	SectionTitle lipgloss.Style
	JobTitle     lipgloss.Style
	CompanyText  lipgloss.Style
	DateText     lipgloss.Style
	BulletText   lipgloss.Style
	Tag          lipgloss.Style
	LinkText     lipgloss.Style
	AsciiStyle   lipgloss.Style
	ContentPad   lipgloss.Style
}

// NewStyles builds all styles from the given theme.
func NewStyles(t Theme) Styles {
	return Styles{
		TabActive:   lipgloss.NewStyle().Background(t.Accent).Foreground(t.Background).Bold(true).Blink(true).PaddingLeft(1).PaddingRight(1),
		TabInactive: lipgloss.NewStyle().Foreground(t.Muted),
		Header:      lipgloss.NewStyle().Foreground(t.Dim),
		Footer:      lipgloss.NewStyle().Foreground(t.Muted),
		SectionTitle: lipgloss.NewStyle().
			Foreground(t.Accent).
			Bold(true).
			MarginTop(1),
		JobTitle:    lipgloss.NewStyle().Foreground(t.Accent).Bold(true),
		CompanyText: lipgloss.NewStyle().Foreground(t.Muted),
		DateText:    lipgloss.NewStyle().Foreground(t.Dim),
		BulletText:  lipgloss.NewStyle().Foreground(t.Foreground),
		Tag: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(t.Muted).
			Foreground(t.Muted).
			Padding(0, 1),
		LinkText:   lipgloss.NewStyle().Foreground(t.Link).Underline(true),
		AsciiStyle: lipgloss.NewStyle().Foreground(t.Dim),
		ContentPad: lipgloss.NewStyle().PaddingLeft(2),
	}
}
