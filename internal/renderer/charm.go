package renderer

import (
	"fmt"
	"log"

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
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	text := m.textInput.Value()

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
		"Enter Cron Expression: \n\n%s\n\n%s\n\n%s",
		m.textInput.View(),
		"(esc/ctrl+c to quit; ctrl+k to increment next at; ctrl+j to decrement next at)",
		m.output,
	) + "\n"
}

func CharmRenderer(cp Parser) {
	p := tea.NewProgram(initialModel(cp))
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Print("\033[H")
}
