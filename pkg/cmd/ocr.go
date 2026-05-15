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

var ocrProcess = requestflag.WithInnerFlags(cli.Command{
	Name:    "process",
	Usage:   "Process a document through Mistral OCR.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "document",
			Usage:    "Document input for OCR.",
			Required: true,
			BodyPath: "document",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Default:  "mistral-ocr-latest",
			BodyPath: "model",
		},
	},
	Action:          handleOCRProcess,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"document": {
		&requestflag.InnerFlag[string]{
			Name:       "document.document-url",
			Usage:      "Data URI with base64-encoded document",
			InnerField: "document_url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "document.type",
			InnerField: "type",
		},
	},
})

func handleOCRProcess(ctx context.Context, cmd *cli.Command) error {
	client := githubcomdedaluslabsdedalussdkgo.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := githubcomdedaluslabsdedalussdkgo.OCRProcessParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.OCR.Process(ctx, params, options...)
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
		Title:          "ocr process",
		Transform:      transform,
	})
}
