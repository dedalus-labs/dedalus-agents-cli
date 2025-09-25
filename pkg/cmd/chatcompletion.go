// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/dedalus-labs/dedalus-cli/pkg/jsonflag"
	"github.com/dedalus-labs/dedalus-sdk-go"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/urfave/cli/v3"
)

var chatCompletionsCreate = cli.Command{
	Name:  "create",
	Usage: "Create a chat completion using the Agent framework.",
	Flags: []cli.Flag{
		&jsonflag.JSONFloatFlag{
			Name: "frequency-penalty",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "frequency_penalty",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "max-tokens",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "max_tokens",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "max-turns",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "max_turns",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "mcp-servers",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "mcp_servers.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+mcp_server",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "mcp_servers.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "model",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "model.name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.name",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name: "model.frequency_penalty",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.frequency_penalty",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "model.logprobs",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "model.logprobs",
				SetValue: true,
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "model.max_completion_tokens",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.max_completion_tokens",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "model.max_tokens",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.max_tokens",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "model.n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.n",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "model.parallel_tool_calls",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "model.parallel_tool_calls",
				SetValue: true,
			},
		},
		&jsonflag.JSONFloatFlag{
			Name: "model.presence_penalty",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.presence_penalty",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "model.seed",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.seed",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "model.service_tier",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.service_tier",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "model.stop",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.stop",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "model.+stop",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.stop.-1",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "model.stream",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "model.stream",
				SetValue: true,
			},
		},
		&jsonflag.JSONFloatFlag{
			Name: "model.temperature",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.temperature",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "model.tool_choice",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.tool_choice",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "model.top_logprobs",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.top_logprobs",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name: "model.top_p",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.top_p",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "model.user",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.user",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+model",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "model.-1",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "n",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "n",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name: "presence-penalty",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "presence_penalty",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "stop",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "stop.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+stop",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "stop.-1",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name: "temperature",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "temperature",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "tool-choice",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "tool_choice",
			},
		},
		&jsonflag.JSONFloatFlag{
			Name: "top-p",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "top_p",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "user",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "user",
			},
		},
	},
	Action:          handleChatCompletionsCreate,
	HideHelpCommand: true,
}

func handleChatCompletionsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParams{}
	stream := cc.client.Chat.Completions.NewStreaming(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
	for stream.Next() {
		fmt.Printf("%s\n", stream.Current().RawJSON())
	}
	return stream.Err()
}
