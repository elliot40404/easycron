package parser

import (
	"fmt"
	"strings"
	"time"

	"github.com/lnquy/cron"
	cp "github.com/robfig/cron/v3"
)

type CronParser struct {
	expr       string
	parser     cp.Schedule
	descriptor *cron.ExpressionDescriptor
}

func NewCronParser() *CronParser {
	return &CronParser{}
}

// Update the cron expression.
func (c *CronParser) SetExpr(expr string) error {
	schedule, err := cp.ParseStandard(expr)
	if err != nil {
		return fmt.Errorf("error parsing cron expression: %w", err)
	}
	descriptor, err := cron.NewDescriptor()
	if err != nil {
		return fmt.Errorf("error creating cron descriptor: %w", err)
	}
	c.expr = expr
	c.descriptor = descriptor
	c.parser = schedule
	return nil
}

// Note: If the length is less than 1, it will return nil.
func (c CronParser) NextInstances(length int) ([]string, error) {
	if length < 1 {
		return nil, nil
	}
	now := time.Now()
	nextTimes := []string{}
	for i := 0; i < length; i++ {
		next := c.parser.Next(now)
		nextTimes = append(nextTimes, next.Format("2006-01-02 15:04:05"))
		now = next
	}
	return nextTimes, nil
}

// Returns a human readable string of the cron expression.
func (c CronParser) HumanReadableStr() (string, error) {
	desc, err := c.descriptor.ToDescription(c.expr, cron.Locale_en)
	if err != nil {
		return "", fmt.Errorf("error creating human readable string: %w", err)
	}
	return desc, nil
}

var HINTS = [...]string{"minute", "hour", "day of month", "month", "day of week"}

func (c CronParser) GetHints(padding, hintIdx int) string {
	if hintIdx > len(HINTS)-1 {
		panic("out of bounds")
	}
	str := strings.Builder{}
	spaces := strings.Repeat(" ", padding)
	str.Grow(len(spaces)*2 + 20)
	str.WriteString(spaces)
	str.WriteString("│\n" + spaces + "│\n")
	hint := HINTS[hintIdx]
	if len(hint) > 8 {
		newPad := padding - len(hint)/2
		if newPad < 0 {
			newPad = 0
		}
		spaces = strings.Repeat(" ", newPad)
	}
	str.WriteString(spaces)
	str.WriteString(hint)
	return str.String()
}
