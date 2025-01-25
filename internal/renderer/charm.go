package renderer

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type model struct {
	parser    Parser
	textInput textinput.Model
	err       error
	output    string
	hints     string
}

func initialModel(parser Parser) model {
	ti := textinput.New()
	ti.Placeholder = "* * * * *"
	ti.Focus()
	ti.Width = 20

	return model{
		parser:    parser,
		textInput: ti,
		err:       nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	isSpace := false
	text := m.textInput.Value()
	items := strings.Split(text, " ")
	inpLen := len(items)
	updateHint := true

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			return m, tea.Quit
		case tea.KeyCtrlK:
			m.parser.IncIter()
		case tea.KeyCtrlJ:
			m.parser.DecIter()
		case tea.KeySpace:
			isSpace = true
		case tea.KeyLeft, tea.KeyRight:
			cpos := m.textInput.Position()
			idx := cpos / 2
			if idx <= 4 {
				padding := 0
				for _, item := range items[0:idx] {
					padding += len(item) + 1
				}
				m.hints = m.parser.GetHints(padding+len(m.textInput.Prompt), idx)
				updateHint = false
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	if inpLen <= 5 && updateHint {
		currElem := items[inpLen-1]
		padding := len(text) - len(currElem)
		m.hints = m.parser.GetHints(padding+len(m.textInput.Prompt), inpLen-1)
	}

	if (text == "" || strings.HasSuffix(text, " ") || len(strings.Split(strings.TrimSpace(text), " ")) >= 5) && isSpace {
		return m, cmd
	}

	text = strings.TrimSpace(text)

	if text != "" {
		err := m.parser.SetExpr(text)
		if err != nil {
			m.output = "invalid cron expression"
			m.textInput, cmd = m.textInput.Update(msg)
			return m, cmd
		}
		m.output = fmt.Sprint(m.parser)
	}

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

func CharmRenderer(cp Parser) {
	p := tea.NewProgram(initialModel(cp), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
