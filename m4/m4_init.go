package m4

// A simple example demonstrating the use of multiple text input components
// from the Bubbles component library.

import (
	"fmt"
	"os"
	"runtime"

	"github.com/charmbracelet/bubbles/key"
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
		RunInitPrompt()
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

func InitModel(bg3Path *string, modsFolderPath *string) Model {
	bg3PathSuggestion := GetBg3PathSuggestion()
	var width int
	if 100 > len(bg3PathSuggestion) {
		width = 100
	} else {
		width = len(bg3PathSuggestion) + 10
	}
	model := Model{width: width}
	model.renderer = lg.DefaultRenderer()
	model.styles = newStyles(model.renderer)
	model.form = huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Where is Baldur's Gate 3 Installed on Your System?").
				Prompt("> ").
				Placeholder(bg3PathSuggestion).
				Suggestions([]string{bg3PathSuggestion}).
				Value(bg3Path),
			huh.NewInput().
				Title("Where will you put your mods?").
				Prompt("> ").Suggestions([]string{"Could be this"}).
				Value(modsFolderPath))).
		WithWidth(model.width).
		WithTheme(huh.ThemeBase()).WithKeyMap(&huh.KeyMap{
		Input: huh.InputKeyMap{
			AcceptSuggestion: key.NewBinding(key.WithKeys("tab"), key.WithHelp("tab", "complete")),
			Prev:             key.NewBinding(key.WithKeys("shift+tab"), key.WithHelp("shift+tab", "back")),
			Next:             key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "next")),
			Submit:           key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "submit")),
		}})
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
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
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

func RunInitPrompt() {
	var bg3Path string
	var modsFolderPath string
	_, err := tea.NewProgram(InitModel(&bg3Path, &modsFolderPath), tea.WithAltScreen()).Run()
	if err != nil {
		fmt.Println("Init prompt encountered and unexpected error", err)
		os.Exit(1)
	}

	// Assign default value if user declined to input.
	if bg3Path == "" {
		bg3Path = GetBg3PathSuggestion()
	}

	config := Config{
		Bg3Path:       bg3Path,
		ModFolderPath: modsFolderPath,
		Mods:          []Mod{},
	}
	config.SaveConfig()
}

// TODO: make this smart enough to figure out the difference between a regular linux machine and the steam deck
func GetBg3PathSuggestion() string {
	switch runtime.GOOS {
	case "linux":
		return "~/deck/.local/share/Steam/steamapps/compatdata/1086940/pfx/drive_C/users/steamuser/appdata/local/Larian Studios"
	default:
		return ""
	}
}
