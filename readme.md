# gscan

`gscan` is a command-line tool that utilizes headless browsing and AI to search GitHub and summarize the results. It incorporates two-factor authentication and can handle various search types on GitHub.

## Features

- **GitHub Login**: Authenticates users using GitHub credentials and OTP verification.
- **Search on GitHub**: Performs searches across different types (code, repositories, issues, etc.) and paginates results.
- **Content Summarization**: Uses Gemini to summarize the HTML content of search results.
- **Screenshot Capability**: Saves screenshots of the results in PNG format.

## Getting Started

### Prerequisites

- Go (version 1.24.0 or later)
- Set up your environment variables:
  - `GITHUB_USER`: Your GitHub username.
  - `GITHUB_PASS`: Your GitHub password.
  - `TOTP_SECRET`: Your TOTP secret for generating OTP.
  - `GEMINI_API_KEY`: Your Gemini API key for summarization.

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/esmaeelnabil/gscan.git
   ```

2. Navigate to the project directory:

   ```bash
   cd gscan
   ```

3. Install dependencies:

   ```bash
   go mod tidy
   ```

### Usage

To perform a search or login, running the command can be done via:

```bash
go run main.go --query "your search query" --type "code" --count 5 --login
```

### Command-Line Flags

- `-query`: The search query on GitHub.
- `-type`: The type of search (options: `code`, `repositories`, `issues`, `pullrequests`, `users`, `commits`).
- `-count`: Number of result pages to fetch (max 5).
- `-login`: Flag to log in and persist the session.
- `-v`: Verbose mode.

## How It Works

1. The user provides GitHub credentials and invokes the application with a query.
2. The application logs in to GitHub, authenticating via the provided credentials and OTP.
3. It performs the search and fetches results.
4. The application summarizes the results using the Gemini service.
5. Optionally, it saves screenshots of the page.

## File Overview

- **engine/chromedp.go**: Contains functions to manage the Chrome context and execute actions.
- **gemini/client.go**: Handles summarization using Gemini API.
- **github/github.go**: Interacts with the GitHub API for login and search functions.
- **lib/utils.go**: Utility functions, such as saving screenshots.
- **main.go**: Entry point for the application, managing command-line interface.
- **otp/totp.go**: Generates TOTP codes for two-factor authentication.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments

- **Chromedp**: For headless Chrome control.
- **Gemini API**: For AI-driven content summarization.
- **pquerna/otp**: For TOTP generation.

---

For further information and contributions, please check the [Wiki](https://github.com/esmaeelnabil/gscan/wiki) or open an issue on the repository.
