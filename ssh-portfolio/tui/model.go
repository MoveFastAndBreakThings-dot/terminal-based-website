package tui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	headerHeight = 1
	tabBarHeight = 2 // tab names + rule
	footerHeight = 1
)

// Model is the bubbletea application model — one instance per SSH session.
type Model struct {
	tabs       []string
	activeTab  int
	viewport   viewport.Model
	theme      Theme
	themeIndex int
	width      int
	height     int
	ready      bool
}

// NewModel creates a fresh model.  width/height come from the SSH PTY.
func NewModel(width, height int) Model {
	m := Model{
		tabs:       []string{"About", "Experience", "Projects", "Skills", "Links"},
		activeTab:  0,
		theme:      Themes[0],
		themeIndex: 0,
		width:      width,
		height:     height,
	}
	if width > 0 && height > 0 {
		vh := vpHeight(height)
		m.viewport = viewport.New(width, vh)
		m.viewport.SetContent(renderActiveTab(m))
		m.ready = true
	}
	return m
}

func vpHeight(totalH int) int {
	vh := totalH - headerHeight - tabBarHeight - footerHeight - 1
	if vh < 1 {
		return 1
	}
	return vh
}

// ── bubbletea interface ───────────────────────────────────────────────────────

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		vh := vpHeight(msg.Height)
		if !m.ready {
			m.viewport = viewport.New(msg.Width, vh)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = vh
		}
		m.viewport.SetContent(renderActiveTab(m))
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		// Tab navigation
		case "tab", "right", "l":
			m.activeTab = (m.activeTab + 1) % len(m.tabs)
			m.viewport.GotoTop()
			m.viewport.SetContent(renderActiveTab(m))
			return m, nil

		case "shift+tab", "left", "h":
			m.activeTab = (m.activeTab - 1 + len(m.tabs)) % len(m.tabs)
			m.viewport.GotoTop()
			m.viewport.SetContent(renderActiveTab(m))
			return m, nil

		// Theme cycling
		case "t":
			m.themeIndex = (m.themeIndex + 1) % len(Themes)
			m.theme = Themes[m.themeIndex]
			m.viewport.SetContent(renderActiveTab(m))
			return m, nil

		// Jump to tab by number
		case "1":
			m = jumpTab(m, 0)
			return m, nil
		case "2":
			m = jumpTab(m, 1)
			return m, nil
		case "3":
			m = jumpTab(m, 2)
			return m, nil
		case "4":
			m = jumpTab(m, 3)
			return m, nil
		case "5":
			m = jumpTab(m, 4)
			return m, nil

		// Quit
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	// Delegate all other messages (scroll keys, mouse, etc.) to the viewport.
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	if !m.ready {
		return "\n  Connecting…"
	}
	s := NewStyles(m.theme)
	return renderHeader(m, s) + "\n" +
		renderTabBar(m, s) + "\n" +
		m.viewport.View() + "\n" +
		renderFooter(m, s)
}

// ── helpers ───────────────────────────────────────────────────────────────────

func jumpTab(m Model, idx int) Model {
	if idx >= 0 && idx < len(m.tabs) {
		m.activeTab = idx
		m.viewport.GotoTop()
		m.viewport.SetContent(renderActiveTab(m))
	}
	return m
}
