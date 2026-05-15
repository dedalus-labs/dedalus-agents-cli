# Dedalus CLI

The official CLI for the [Dedalus REST API](https://docs.dedaluslabs.ai).

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

To test or install the CLI locally, you need [Go](https://go.dev/doc/install) version 1.22 or later installed.

```sh
go install 'github.com/stainless-sdks/dedalus-sdk-cli/cmd/dedalus-sdk@latest'
```

Once you have run `go install`, the binary is placed in your Go bin directory:

- **Default location**: `$HOME/go/bin` (or `$GOPATH/bin` if GOPATH is set)
- **Check your path**: Run `go env GOPATH` to see the base directory

If commands aren't found after installation, add the Go bin directory to your PATH:

```sh
# Add to your shell profile (.zshrc, .bashrc, etc.)
export PATH="$PATH:$(go env GOPATH)/bin"
```

### Running Locally

After cloning the git repository for this project, you can use the
`scripts/run` script to run the tool locally:

```sh
./scripts/run args...
```

## Usage

The CLI follows a resource-based command structure:

```sh
dedalus-sdk [resource] <command> [flags...]
```

```sh
dedalus-sdk chat:completions create \
  --api-key 'My API Key' \
  --model openai/gpt-5-nano \
  --message '{content: You are Stephen Dedalus. Respond in morose Joycean malaise., role: system}' \
  --message "{content: 'Hello, how are you today?', role: user}"
```

For details about specific commands, use the `--help` flag.

### Environment variables

| Environment variable     | Description                                                          | Required | Default value                 |
| ------------------------ | -------------------------------------------------------------------- | -------- | ----------------------------- |
| `DEDALUS_API_KEY`        | API key for Bearer token authentication.                             | no       | `null`                        |
| `DEDALUS_X_API_KEY`      | API key for X-API-Key header authentication.                         | no       | `null`                        |
| `DEDALUS_AS_URL`         | MCP Authorization Server URL                                         | no       | `"https://as.dedaluslabs.ai"` |
| `DEDALUS_ORG_ID`         | Organization ID for request scoping.                                 | no       | `null`                        |
| `DEDALUS_PROVIDER`       | Provider name for BYOK mode (e.g., 'google', 'openai', 'anthropic'). | no       | `null`                        |
| `DEDALUS_PROVIDER_KEY`   | Provider API key for BYOK mode.                                      | no       | `null`                        |
| `DEDALUS_PROVIDER_MODEL` | Model identifier for BYOK provider.                                  | no       | `null`                        |

### Global flags

- `--api-key` - API key for Bearer token authentication. (can also be set with `DEDALUS_API_KEY` env var)
- `--x-api-key` - API key for X-API-Key header authentication. (can also be set with `DEDALUS_X_API_KEY` env var)
- `--as-base-url` - MCP Authorization Server URL (can also be set with `DEDALUS_AS_URL` env var)
- `--dedalus-org-id` - Organization ID for request scoping. (can also be set with `DEDALUS_ORG_ID` env var)
- `--provider` - Provider name for BYOK mode (e.g., 'google', 'openai', 'anthropic'). (can also be set with `DEDALUS_PROVIDER` env var)
- `--provider-key` - Provider API key for BYOK mode. (can also be set with `DEDALUS_PROVIDER_KEY` env var)
- `--provider-model` - Model identifier for BYOK provider. (can also be set with `DEDALUS_PROVIDER_MODEL` env var)
- `--help` - Show command line usage
- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
- `--base-url` - Use a custom API backend URL
- `--format` - Change the output format (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--format-error` - Change the output format for errors (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--transform` - Transform the data output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
- `--transform-error` - Transform the error output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)

### Passing files as arguments

To pass files to your API, you can use the `@myfile.ext` syntax:

```bash
dedalus-sdk <command> --arg @abe.jpg
```

Files can also be passed inside JSON or YAML blobs:

```bash
dedalus-sdk <command> --arg '{image: "@abe.jpg"}'
# Equivalent:
dedalus-sdk <command> <<YAML
arg:
  image: "@abe.jpg"
YAML
```

If you need to pass a string literal that begins with an `@` sign, you can
escape the `@` sign to avoid accidentally passing a file.

```bash
dedalus-sdk <command> --username '\@abe'
```

#### Explicit encoding

For JSON endpoints, the CLI tool does filetype sniffing to determine whether the
file contents should be sent as a string literal (for plain text files) or as a
base64-encoded string literal (for binary files). If you need to explicitly send
the file as either plain text or base64-encoded data, you can use
`@file://myfile.txt` (for string encoding) or `@data://myfile.dat` (for
base64-encoding). Note that absolute paths will begin with `@file://` or
`@data://`, followed by a third `/` (for example, `@file:///tmp/file.txt`).

```bash
dedalus-sdk <command> --arg @data://file.txt
```

## Linking different Go SDK versions

You can link the CLI against a different version of the Dedalus Go SDK
for development purposes using the `./scripts/link` script.

To link to a specific version from a repository (version can be a branch,
git tag, or commit hash):

```bash
./scripts/link github.com/org/repo@version
```

To link to a local copy of the SDK:

```bash
./scripts/link ../path/to/githubcomdedaluslabsdedalussdkgo-go
```

If you run the link script without any arguments, it will default to `../githubcomdedaluslabsdedalussdkgo-go`.
