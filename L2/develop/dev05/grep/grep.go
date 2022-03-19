package grep

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

type flags struct {
	a     *int
	b     *int
	c     *int
	count *bool
	i     *bool
	v     *bool
	f     *bool
	n     *bool
}

// Grep structure
type Grep struct {
	flags
	data []string
}

// NewGrep grep's structure constructor
func NewGrep() *Grep {
	return &Grep{
		flags: flags{
			a:     flag.Int("A", 0, "print +N lines after match"),
			b:     flag.Int("B", 0, "print +N lines before match"),
			c:     flag.Int("C", 0, "(A+B) print ±N lines around the match"),
			count: flag.Bool("c", false, "number of lines"),
			i:     flag.Bool("i", false, "ignore case"),
			v:     flag.Bool("v", false, "instead of match, exclude"),
			f:     flag.Bool("F", false, "exact string match, not a pattern"),
			n:     flag.Bool("n", false, "print line number"),
		},
		data: make([]string, 0),
	}
}

// Execute runs the utility
func (g *Grep) Execute() error {
	flag.Parse()

	pattern := flag.Arg(0)
	if len(pattern) < 1 {
		return errors.New("grep: missing pattern")
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		g.data = append(g.data, scanner.Text())
	}

	if *g.c > 0 {
		*g.a = *g.c
		*g.b = *g.c
	}

	counter, last := 0, 0
	for n, line := range g.data {
		ok := strings.Contains(line, pattern) || (*g.i && strings.Contains(strings.ToLower(line), strings.ToLower(pattern)))

		if *g.f {
			if pattern != line {
				continue
			}
		}

		if (ok && !*g.v) || (!ok && *g.v) {
			counter++

			for *g.b > 0 {
				if n-*g.b < 0 {
					*g.b--
					continue
				}

				if *g.n {
					fmt.Printf("%d-%s\n", n-*g.b+1, g.data[n-*g.b])
				} else {
					fmt.Println(g.data[n-*g.b])
				}
				*g.b--
			}

			if *g.n {
				fmt.Printf("%d:%s\n", n+1, line)
			} else {
				fmt.Println(line)
			}

			last = n
		}
	}

	for i := 0; i < *g.a; i++ {
		if last+1+i >= len(g.data) {
			continue
		}

		if *g.n {
			fmt.Printf("%d-%s\n", last+1+i+1, g.data[last+1+i])
		} else {
			fmt.Println(g.data[last+1+i])
		}
	}

	if *g.count {
		fmt.Println(counter)
	}

	return nil
}
