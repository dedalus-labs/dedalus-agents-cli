// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/urfave/cli/v3"
)

var rootGet = cli.Command{
	Name:            "get",
	Usage:           "Root",
	Flags:           []cli.Flag{},
	Action:          handleRootGet,
	HideHelpCommand: true,
}

func handleRootGet(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Root.Get(
		context.TODO(),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("root get", string(res), format)
}
