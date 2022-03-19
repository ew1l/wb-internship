package wget

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

// Wget structure
type Wget struct {
	url string
	c   *colly.Collector
}

// NewWget wget's structure constructor
func NewWget(depth int) *Wget {
	return &Wget{
		url: "",
		c:   colly.NewCollector(colly.MaxDepth(depth)),
	}
}

// Execute runs the utility
func (w *Wget) Execute() error {
	flag.Parse()

	if len(flag.Args()) < 1 {
		return errors.New("wget: missing URL")
	}

	w.url = strings.TrimRight(flag.Arg(0), "/")
	parsed, err := url.Parse(w.url)
	if err != nil {
		return err
	}

	if parsed.Scheme == "" {
		return errors.New("wget: missing URL scheme (example: https://)")
	}

	w.c.AllowedDomains = append(w.c.AllowedDomains, parsed.Host)

	dir := parsed.Host
	err = os.Mkdir(dir, 0600)
	if err != nil {
		return err
	}

	w.c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		w.c.Visit(e.Request.AbsoluteURL(e.Attr("href")))
	})

	w.c.OnRequest(func(r *colly.Request) {
		log.Println(r.URL.String())
	})

	w.c.OnResponse(func(r *colly.Response) {
		resp, err := http.Get(r.Request.URL.String())
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		content, _ := ioutil.ReadAll(resp.Body)

		parsed, _ := url.Parse(r.Request.URL.String())
		file := dir + "/" + strings.ReplaceAll(parsed.Host+parsed.Path, "/", ".")

		os.Create(file)

		err = ioutil.WriteFile(file, content, 0600)
		if err != nil {
			log.Println(err)
		}
	})

	w.c.Visit(w.url)

	return nil
}
