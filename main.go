package main

import (
	"flag"
	"fmt"
	"gscan/engine"
	"gscan/gemini"
	"gscan/github"
	"os"

	"github.com/chromedp/chromedp"
)

var prompt string = ""

func main() {
	query := flag.String("query", "", "Github Search Query")
	searchType := flag.String("type", "code", "Github Search Query type, code || repositories || issues || pullrequests || users || commits")
	count := flag.Int("count", 1, "Github Search Query Code Pages count, max 5")
	loginIsNeeded := false
	flag.BoolFunc("login", "Will Login and presist the user session of the engine for later usage", func(s string) error {
		loginIsNeeded = true
		return nil
	})
	isVerbose := false
	flag.BoolFunc("v", "verbose messages on each step", func(s string) error {
		isVerbose = true
		return nil
	})
	flag.Parse()

	allocCtx, cancelAllocator := engine.GetContext()
	defer cancelAllocator()
	browserContext, cancelBrowserContext := chromedp.NewContext(allocCtx)
	defer cancelBrowserContext()

	if loginIsNeeded {
		github.Login(isVerbose, browserContext)
		fmt.Println("Successfully Logged in ..")
		os.Exit(0)
	}

	if len(*query) > 0 {
		if *searchType == "code" {
			for page := 1; page <= *count; page++ {
				html, err := github.GithubSearch(browserContext, *query, *searchType, page)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println("Thinking ... ")
				gemini.Summarize(fmt.Sprintf("This Html source is the search result for a search query in github.com search, convert it into md format or a format that would be easially reradable in the unix terminals, the query was (%s) and the content is : %s", *query, html))
			}
		} else {

			html, err := github.GithubSearch(browserContext, *query, *searchType, 1)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Thinking ... ")
			gemini.Summarize(fmt.Sprintf(prompt, query, html))

		}
	}
}
