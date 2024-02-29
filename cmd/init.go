package cmd

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	lg "github.com/charmbracelet/lipgloss"
	cobra "github.com/spf13/cobra"
)

func init() {
	rootCommand.AddCommand(initCommand)
}

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize Mordenkainen’s Magnificent Mod Manager",
	Long:  `Run this when setting up on a new system (or wanting the wizard to help you update your core configs)`,
	Run: func(cmd *cobra.Command, args []string) {
		InitPrompt()
	},
}

type Styles struct {
	Base,
	Header,
	Form lg.Style
}

type Model struct {
	renderer *lg.Renderer
	styles   *Styles
	form     *huh.Form
	width    int
}

func newStyles(renderer *lg.Renderer) *Styles {
	styles := Styles{}
	styles.Base = renderer.
		NewStyle().
		Padding(1).
		BorderStyle(lg.RoundedBorder()).
		BorderForeground(lg.Color("#e40712"))
	styles.Header = renderer.
		NewStyle().
		Foreground(lg.Color("#e40712"))
	styles.Form = renderer.
		NewStyle().
		PaddingTop(1)
	return &styles
}

func InitModel(bg3Path string, modsFolderPath string) Model {
	model := Model{width: 100}
	model.renderer = lg.DefaultRenderer()
	model.styles = newStyles(model.renderer)
	model.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Where is Baldur's Gate 3 Installed on Your System?").
				Prompt("> ").Placeholder("").
				Value(&bg3Path),

			huh.NewInput().
				Title("Where will you put your mods?").
				Prompt("> ").
				Value(&modsFolderPath),
		),
	).WithWidth(model.width).WithTheme(huh.ThemeBase())
	return model
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) View() string {
	header := m.styles.Header.Render("Welcome To Mordenkainen’s Magnificent Mod Manager!")
	finalForm := m.styles.Form.Render(m.form.View())
	return m.styles.Base.Render(lg.JoinVertical(lg.Top, header, finalForm))
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Update the form in sync with the app.
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	// Quit when the form is done.
	if m.form.State == huh.StateCompleted {
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func InitPrompt() {
	var bg3Path string
	var modsFolderPath string
	_, err := tea.NewProgram(InitModel(bg3Path, modsFolderPath), tea.WithAltScreen()).Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
}
