package engine

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"
)

func LogStep(step string) chromedp.Action {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		fmt.Println(step)
		return nil
	})
}

func GetContext() (context.Context, context.CancelFunc) {
	// Get the user's home directory.
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	// Define the profile directory inside the home folder.
	profileDir := filepath.Join(home, ".gscan_profile")
	// Check if the directory exists.
	if _, err := os.Stat(profileDir); os.IsNotExist(err) {
		log.Println("You might need to login first\nUse : gscan -login")
	}
	// creating chrome context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Headless,
		chromedp.NoSandbox,
		chromedp.UserDataDir(profileDir),
		chromedp.Flag("enable-automation", false),
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("start-maximized", true),
	)

	allocCtx, cancelAllocator := chromedp.NewExecAllocator(context.Background(), opts...)

	return allocCtx, cancelAllocator
}

func GetLoginTasks(username, password, otp string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("https://github.com/login"),
		chromedp.WaitVisible(`input[name="login"]`),
		chromedp.SendKeys(`input[name="login"]`, username),
		chromedp.WaitVisible(`input[name="password"]`),
		chromedp.SendKeys(`input[name="password"]`, password),
		chromedp.WaitVisible(`input[name="commit"][value="Sign in"]`),
		chromedp.Click(`input[name="commit"][value="Sign in"]`),

		// Wait for the 2FA recovery link and click it
		chromedp.WaitVisible(`a[data-test-selector="totp-app-link"]`),
		chromedp.Click(`a[data-test-selector="totp-app-link"]`),
		// Wait for the recovery code input field
		chromedp.WaitVisible(`input[name="app_otp"]`),
		LogStep("- OTP Visible ✅"),
		chromedp.SendKeys(`input[name="app_otp"]`, otp),
		LogStep("- OTP Done ✅"),
		chromedp.Sleep(1 * time.Second),
	}
}

func GetSearchTasks(query, searchType string, page int) chromedp.Tasks {
	return chromedp.Tasks{
		LogStep("- Searching for : " + query),
		chromedp.Navigate(fmt.Sprintf("https://github.com/search?q=%s&type=%s&p=%d", query, searchType, page)),
	}
}
