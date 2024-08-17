# Redmine Issues Fetcher

This Go application fetches and lists issues from a Redmine project. It queries the Redmine server via its REST API, fetches the issues in JSON format, and prints them out in a readable format.

## Prerequisites

Before you can use this tool, you'll need the following:

1. **Go**: Make sure you have Go installed on your system. You can download it from [golang.org](https://golang.org/).

2. **Redmine API Key**: Obtain your Redmine API key from your Redmine account settings.

## Installation

1. Clone the repository or download the source code.

    ```sh
    git clone <repository-url>
    cd <repository-directory>
    ```

2. Build the executable.

    ```sh
    go build -o redmine-issues-fetcher
    ```

## Usage

1. Set the `REDMINE_API_KEY` environment variable with your Redmine API key.

    ```sh
    export REDMINE_API_KEY=your_api_key_here
    ```

2. Run the application with the Redmine URL as an argument.

    ```sh
    ./redmine-issues-fetcher https://your-redmine-url/
    ```

## Example

```sh
export REDMINE_API_KEY=abcdefgh12345678
./redmine-issues-fetcher https://redmine.example.com/
```

## Disclaimer

This tool disables TLS certificate verification, which might expose you to security risks. Ensure that you understand the implications and use it in controlled environments only.

---

Feel free to modify and improve this README to better suit your needs!