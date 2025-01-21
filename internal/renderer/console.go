package renderer

import (
	"fmt"
	"os"
)

func ConsoleRenderer(cp Parser) {
	err := cp.Validate()
	if err != nil {
		fmt.Println("Invalid cron expression")
		os.Exit(1)
	}
	fmt.Print(cp)
}
