package main

import (
	"flag"
	"fmt"
	"gscan/ai"
	"gscan/engine"
	"gscan/github"
	"os"

	"github.com/chromedp/chromedp"
)

var (
	systemMessage string = "You will recieve a search query and Html source for the result of the search query in github.com search, extract from it the outcome of the search only then list it with the code/repositorie/issue/etc under the bullet point, so the user will excpect a list of results that is easy to read without side menues or anything else other than the result for the query, the output should be easialy reradable in any unix terminal, so do not wrap or use any formatting for your input or explain anything."
	prompt        string = ""
)

func main() {
	query := flag.String("query", "", "Github Search Query")
	searchType := flag.String("type", "code", "Github Search Query type, code || repositories || issues || pullrequests || users || commits")
	count := flag.Int("page-count", 1, "Github Search Query Code Pages count, max 5")
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

	var agent ai.AiAgent = ai.OpenAiAgent{}

	if len(*query) > 0 {
		if *searchType == "code" {
			for page := 1; page <= *count; page++ {
				html, err := github.GithubSearch(browserContext, *query, *searchType, page)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Println("Thinking ... ")
				prompt := fmt.Sprintf("Query : %s\n Html source result : %s", *query, html)
				fmt.Println(agent.Process(prompt, "Model isn't used for now", systemMessage))
			}
		} else {

			html, err := github.GithubSearch(browserContext, *query, *searchType, 1)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println("Thinking ... ")

			prompt := fmt.Sprintf("Query : %s\n Html source result : %s", *query, html)

			fmt.Println(agent.Process(prompt, "Model isn't used for now", systemMessage))
		}
	}
}
