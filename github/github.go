package github

import (
	"context"
	"fmt"
	"gscan/engine"
	"gscan/otp"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func GithubSearch(ctx context.Context, query, searchType string, page int) (string, error) {
	searchTasks := engine.GetSearchTasks(query, searchType, page)

	var htmlContent string

	appendedScreenShotTasks := append(
		searchTasks,
		chromedp.Tasks{
			chromedp.OuterHTML("html", &htmlContent),
		},
	)

	if err := chromedp.Run(ctx, appendedScreenShotTasks); err != nil {
		fmt.Println(err)
		return "", err
	}

	return htmlContent, nil
}

func Login(isVerbose bool, ctx context.Context) {
	username := os.Getenv("GITHUB_USER")
	password := os.Getenv("GITHUB_PASS")
	TwoFASecretKey := os.Getenv("TOTP_SECRET")

	if len(TwoFASecretKey) == 0 {
		log.Fatalf("Make sure you added TOTP_SECRET to your env variables")
		os.Exit(1)
	}

	if len(username) == 0 || len(password) == 0 {
		fmt.Println("Make sure you added GITHUB_USER and GITHUB_PASS to your env variables")
		os.Exit(1)
	}

	otp := otp.GenerateTotp(TwoFASecretKey)

	if isVerbose {
		fmt.Println("- T-OTP: ", otp)
	}

	loginTasks := engine.GetLoginTasks(username, password, otp)

	if err := chromedp.Run(ctx, loginTasks); err != nil {
		fmt.Println(err)
		return
	}
}
