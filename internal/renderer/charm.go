//go:build !tview

package renderer

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
	Charm  struct{}
	model  struct {
		parser    Parser
		textInput textinput.Model
		err       error
		output    string
		hints     string
	}
)

const textWidth = 20

func NewRenderer() Renderer {
	return Charm{}
}

func initialModel(parser Parser) model {
	ti := textinput.New()
	ti.Placeholder = "* * * * *"
	ti.Focus()
	ti.Width = textWidth
	return model{
		parser:    parser,
		textInput: ti,
		err:       nil,
	}
}

func (model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	isSpace := false

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			return m, tea.Quit
		case tea.KeyUp:
			m.parser.IncIter()
		case tea.KeyDown:
			m.parser.DecIter()
		case tea.KeySpace:
			isSpace = true
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	text := m.textInput.Value()
	updateHint(&m, text)

	if !validInput(text, isSpace) {
		return m, cmd
	}

	updateOutput(&m, text)

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"Enter Cron Expression: \n\n%s\n%s\n\n%s",
		m.textInput.View(),
		m.hints,
		m.output,
	) + "\n"
}

func (Charm) Render(cp Parser) error {
	p := tea.NewProgram(initialModel(cp), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		return err
	}
	return nil
}

func updateHint(m *model, text string) {
	if text == "" {
		m.hints = m.parser.GetHints(len(m.textInput.Prompt), 0)
		return
	}
	items := strings.Split(text, " ")
	indexes := findRuneIndexes(text, ' ')
	cpos := m.textInput.Position()
	var idx, padding int
	for i, val := range indexes {
		if cpos <= val {
			idx = i
			break
		}
		if cpos > indexes[len(indexes)-1] && idx < 4 {
			idx = i + 1
		}
	}
	for _, item := range items[0:idx] {
		padding += len(item) + 1
	}
	m.hints = m.parser.GetHints(padding+len(m.textInput.Prompt), idx)
}

func validInput(text string, isSpace bool) bool {
	isEmptyOrEndsWithSpace := text == "" || strings.HasSuffix(text, " ")
	wordCount := len(strings.Split(strings.TrimSpace(text), " "))
	if (isEmptyOrEndsWithSpace || wordCount >= 5) && isSpace {
		return false
	}
	return true
}

func updateOutput(m *model, text string) {
	text = strings.TrimSpace(text)
	if text == "" {
		return
	}
	err := m.parser.SetExpr(text)
	if err != nil {
		m.output = "invalid cron expression"
		return
	}
	m.output = fmt.Sprint(m.parser)
}

func findRuneIndexes(str string, search rune) []int {
	var indexes []int
	for i, r := range str {
		if r == search {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
