// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/urfave/cli/v3"
)

var healthCheck = cli.Command{
	Name:            "check",
	Usage:           "Simple health check.",
	Flags:           []cli.Flag{},
	Action:          handleHealthCheck,
	HideHelpCommand: true,
}

func handleHealthCheck(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Health.Check(
		context.TODO(),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("health check", string(res), format)
}
