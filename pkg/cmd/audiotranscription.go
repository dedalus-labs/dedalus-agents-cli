// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/dedalus-labs/dedalus-sdk-go"
	"github.com/dedalus-labs/dedalus-sdk-go/option"
	"github.com/stainless-sdks/dedalus-sdk-cli/internal/apiquery"
	"github.com/stainless-sdks/dedalus-sdk-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var audioTranscriptionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Transcribe audio into text.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "file",
			Required:  true,
			BodyPath:  "file",
			FileInput: true,
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Required: true,
			BodyPath: "model",
		},
		&requestflag.Flag[*string]{
			Name:     "language",
			BodyPath: "language",
		},
		&requestflag.Flag[*string]{
			Name:     "prompt",
			BodyPath: "prompt",
		},
		&requestflag.Flag[*string]{
			Name:     "response-format",
			BodyPath: "response_format",
		},
		&requestflag.Flag[*float64]{
			Name:     "temperature",
			BodyPath: "temperature",
		},
	},
	Action:          handleAudioTranscriptionsCreate,
	HideHelpCommand: true,
}

func handleAudioTranscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := githubcomdedaluslabsdedalussdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		MultipartFormEncoded,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomdedaluslabsdedalussdkgo.AudioTranscriptionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Audio.Transcriptions.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "audio:transcriptions create",
		Transform:      transform,
	})
}
