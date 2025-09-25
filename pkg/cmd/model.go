// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/urfave/cli/v3"
)

var modelsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get information about a specific model.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "model-id",
		},
	},
	Action:          handleModelsRetrieve,
	HideHelpCommand: true,
}

var modelsList = cli.Command{
	Name:            "list",
	Usage:           "List available models.",
	Flags:           []cli.Flag{},
	Action:          handleModelsList,
	HideHelpCommand: true,
}

func handleModelsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Models.Get(
		context.TODO(),
		cmd.Value("model-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("models retrieve", string(res), format)
}

func handleModelsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Models.List(
		context.TODO(),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("models list", string(res), format)
}
