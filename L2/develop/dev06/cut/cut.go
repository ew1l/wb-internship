package cut

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type flags struct {
	f *int
	d *string
	s *bool
}

// Cut structure
type Cut struct {
	flags
	data []string
}

// NewCut cut's structure constructor
func NewCut() *Cut {
	return &Cut{
		flags: flags{
			f: flag.Int("f", 0, "select fields (columns)"),
			d: flag.String("d", "\t", "use a different delimiter"),
			s: flag.Bool("s", false, "only strings with delimiter"),
		},
		data: make([]string, 0),
	}
}

// Execute runs the utility
func (c *Cut) Execute() error {
	flag.Parse()

	if *c.f < 1 {
		return errors.New("cut: you must specify a field")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		c.data = append(c.data, scanner.Text())
	}

	for _, line := range c.data {
		if strings.Contains(line, *c.d) {
			dline := strings.Split(line, *c.d)

			if *c.f <= len(dline) {
				fmt.Println(dline[*c.f-1])
			} else {
				fmt.Println()
			}
		} else {
			if !*c.s {
				fmt.Println(line)
				continue
			}
			break
		}
	}

	return nil
}
