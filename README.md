# Dedalus CLI

The official CLI for the [Dedalus REST API](https://docs.dedaluslabs.ai).

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

<!-- x-release-please-start-version -->

```sh
go install 'github.com/dedalus-labs/dedalus-cli/cmd/dedalus-sdk@latest'
```

### Running Locally

<!-- x-release-please-start-version -->

```sh
go run cmd/dedalus-sdk/main.go
```

<!-- x-release-please-end -->

## Usage

The CLI follows a resource-based command structure:

```sh
dedalus-sdk [resource] [command] [flags]
```

```sh
dedalus-sdk chat:completions create \
  --model openai/gpt-5
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
