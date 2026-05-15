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

var responsesCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a response using the OpenAI Responses API.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[*bool]{
			Name:     "background",
			Usage:    "Whether to run the model response in the background.\n[Learn more](https://platform.openai.com/docs/guides/background).",
			BodyPath: "background",
		},
		&requestflag.Flag[any]{
			Name:     "conversation",
			Usage:    "Conversation that this response belongs to. Items from this conversation are prepended to the input items, and output items from this response are automatically added after completion.",
			BodyPath: "conversation",
		},
		&requestflag.Flag[any]{
			Name:     "credentials",
			Usage:    "Credentials for MCP server authentication. Each credential is matched to servers by connection name.",
			BodyPath: "credentials",
		},
		&requestflag.Flag[*float64]{
			Name:     "frequency-penalty",
			Usage:    "Penalizes new tokens based on their frequency in the text so far.",
			BodyPath: "frequency_penalty",
		},
		&requestflag.Flag[any]{
			Name:     "include",
			Usage:    "Specify additional output data to include in the model response. Currently\nsupported values are:\n- `web_search_call.action.sources`: Include the sources of the web search tool call.\n- `code_interpreter_call.outputs`: Includes the outputs of python code execution\n  in code interpreter tool call items.\n- `computer_call_output.output.image_url`: Include image urls from the computer call output.\n- `file_search_call.results`: Include the search results of\n  the file search tool call.\n- `message.input_image.image_url`: Include image urls from the input message.\n- `message.output_text.logprobs`: Include logprobs with assistant messages.\n- `reasoning.encrypted_content`: Includes an encrypted version of reasoning\n  tokens in reasoning item outputs. This enables reasoning items to be used in\n  multi-turn conversations when using the Responses API statelessly (like\n  when the `store` parameter is set to `false`, or when an organization is\n  enrolled in the zero data retention program).",
			BodyPath: "include",
		},
		&requestflag.Flag[any]{
			Name:     "input",
			Usage:    "Text, image, or file inputs to the model, used to generate a response.\n\nLearn more:\n- [Text inputs and outputs](https://platform.openai.com/docs/guides/text)\n- [Image inputs](https://platform.openai.com/docs/guides/images)\n- [File inputs](https://platform.openai.com/docs/guides/pdf-files)\n- [Conversation state](https://platform.openai.com/docs/guides/conversation-state)\n- [Function calling](https://platform.openai.com/docs/guides/function-calling)",
			BodyPath: "input",
		},
		&requestflag.Flag[any]{
			Name:     "instructions",
			Usage:    "A system (or developer) message inserted into the model's context.\n\nWhen using along with `previous_response_id`, the instructions from a previous\nresponse will not be carried over to the next response. This makes it simple\nto swap out system (or developer) messages in new responses.",
			BodyPath: "instructions",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-output-tokens",
			Usage:    "An upper bound for the number of tokens that can be generated for a response, including visible output tokens and [reasoning tokens](https://platform.openai.com/docs/guides/reasoning).",
			BodyPath: "max_output_tokens",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-tool-calls",
			Usage:    "The maximum number of total calls to built-in tools that can be processed in a response. This maximum number applies across all built-in tool calls, not per individual tool. Any further attempts to call a tool by the model will be ignored.",
			BodyPath: "max_tool_calls",
		},
		&requestflag.Flag[any]{
			Name:     "mcp-servers",
			Usage:    "MCP server identifiers. Accepts marketplace slugs, URLs, or MCPServerSpec objects. MCP tools are executed server-side and billed separately.",
			BodyPath: "mcp_servers",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Set of up to 16 key-value string pairs that can be attached to the response for structured metadata and later querying via the API or dashboard.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "model",
			Usage:    "Model ID used to generate the response, like `gpt-4o` or `o3`. OpenAI\noffers a wide range of models with different capabilities, performance\ncharacteristics, and price points. Refer to the [model guide](https://platform.openai.com/docs/models)\nto browse and compare available models.",
			BodyPath: "model",
		},
		&requestflag.Flag[*bool]{
			Name:     "parallel-tool-calls",
			Usage:    "Whether to allow the model to run tool calls in parallel.",
			BodyPath: "parallel_tool_calls",
		},
		&requestflag.Flag[*float64]{
			Name:     "presence-penalty",
			Usage:    "Penalizes new tokens based on whether they appear in the text so far.",
			BodyPath: "presence_penalty",
		},
		&requestflag.Flag[*string]{
			Name:     "previous-response-id",
			Usage:    "Unique ID of the previous response to continue from when creating multi-turn conversations. Cannot be used together with `conversation`.",
			BodyPath: "previous_response_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "prompt",
			Usage:    "Stored prompt template reference (BYOK).",
			BodyPath: "prompt",
		},
		&requestflag.Flag[*string]{
			Name:     "prompt-cache-key",
			Usage:    "Used by OpenAI to cache responses for similar requests to optimize your cache hit rates. Replaces the `user` field. [Learn more](https://platform.openai.com/docs/guides/prompt-caching).",
			BodyPath: "prompt_cache_key",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "reasoning",
			BodyPath: "reasoning",
		},
		&requestflag.Flag[*string]{
			Name:     "safety-identifier",
			Usage:    "A stable identifier used to help detect users of your application that may be violating OpenAI's usage policies.\nThe IDs should be a string that uniquely identifies each user. We recommend hashing their username or email address, in order to avoid sending us any identifying information. [Learn more](https://platform.openai.com/docs/guides/safety-best-practices#safety-identifiers).",
			BodyPath: "safety_identifier",
		},
		&requestflag.Flag[*string]{
			Name:     "service-tier",
			Usage:    "Specifies the processing type used for serving the request.\n  - If set to 'auto', then the request will be processed with the service tier configured in the Project settings. Unless otherwise configured, the Project will use 'default'.\n  - If set to 'default', then the request will be processed with the standard pricing and performance for the selected model.\n  - If set to '[flex](https://platform.openai.com/docs/guides/flex-processing)' or '[priority](https://openai.com/api-priority-processing/)', then the request will be processed with the corresponding service tier.\n  - When not set, the default behavior is 'auto'.\n\n  When the `service_tier` parameter is set, the response body will include the `service_tier` value based on the processing mode actually used to serve the request. This response value may be different from the value set in the parameter.",
			BodyPath: "service_tier",
		},
		&requestflag.Flag[*bool]{
			Name:     "store",
			Usage:    "Whether to store the generated response for later retrieval via the Responses API.",
			BodyPath: "store",
		},
		&requestflag.Flag[bool]{
			Name:     "stream",
			Usage:    "If set to true, the model response data will be streamed to the client\nas it is generated using [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format).\nSee the [Streaming section below](https://platform.openai.com/docs/api-reference/responses-streaming)\nfor more information.",
			Default:  false,
			BodyPath: "stream",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "stream-options",
			BodyPath: "stream_options",
		},
		&requestflag.Flag[*float64]{
			Name:     "temperature",
			Usage:    "What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.\nWe generally recommend altering this or `top_p` but not both.",
			BodyPath: "temperature",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "text",
			BodyPath: "text",
		},
		&requestflag.Flag[any]{
			Name:     "tool-choice",
			Usage:    "How the model should select which tool (or tools) to use when generating\na response. See the `tools` parameter to see how to specify which tools\nthe model can call.",
			BodyPath: "tool_choice",
		},
		&requestflag.Flag[any]{
			Name:     "tool",
			Usage:    "An array of tools the model may call while generating a response. You\ncan specify which tool to use by setting the `tool_choice` parameter.\n\nWe support the following categories of tools:\n- **Built-in tools**: Tools that are provided by OpenAI that extend the\n  model's capabilities, like [web search](https://platform.openai.com/docs/guides/tools-web-search)\n  or [file search](https://platform.openai.com/docs/guides/tools-file-search). Learn more about\n  [built-in tools](https://platform.openai.com/docs/guides/tools).\n- **MCP Tools**: Integrations with third-party systems via custom MCP servers\n  or predefined connectors such as Google Drive and SharePoint. Learn more about\n  [MCP Tools](https://platform.openai.com/docs/guides/tools-connectors-mcp).\n- **Function calls (custom tools)**: Functions that are defined by you,\n  enabling the model to call your own code with strongly typed arguments\n  and outputs. Learn more about\n  [function calling](https://platform.openai.com/docs/guides/function-calling). You can also use\n  custom tools to call your own code.",
			BodyPath: "tools",
		},
		&requestflag.Flag[*int64]{
			Name:     "top-logprobs",
			Usage:    "An integer between 0 and 20 specifying the number of most likely tokens to\nreturn at each token position, each with an associated log probability.",
			BodyPath: "top_logprobs",
		},
		&requestflag.Flag[*float64]{
			Name:     "top-p",
			Usage:    "An alternative to sampling with temperature, called nucleus sampling,\nwhere the model considers the results of the tokens with top_p probability\nmass. So 0.1 means only the tokens comprising the top 10% probability mass\nare considered.\n\nWe generally recommend altering this or `temperature` but not both.",
			BodyPath: "top_p",
		},
		&requestflag.Flag[*string]{
			Name:     "truncation",
			Usage:    "The truncation strategy to use for the model response.\n- `auto`: If the input to this Response exceeds\n  the model's context window size, the model will truncate the\n  response to fit the context window by dropping items from the beginning of the conversation.\n- `disabled` (default): If the input size will exceed the context window\n  size for a model, the request will fail with a 400 error.",
			BodyPath: "truncation",
		},
		&requestflag.Flag[*string]{
			Name:     "user",
			Usage:    "This field is being replaced by `safety_identifier` and `prompt_cache_key`. Use `prompt_cache_key` instead to maintain caching optimizations.\nA stable identifier for your end-users.\nUsed to boost cache hit rates by better bucketing similar requests and  to help OpenAI detect and prevent abuse. [Learn more](https://platform.openai.com/docs/guides/safety-best-practices#safety-identifiers).",
			BodyPath: "user",
		},
	},
	Action:          handleResponsesCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"prompt": {
		&requestflag.InnerFlag[string]{
			Name:       "prompt.id",
			Usage:      "Identifier of the stored prompt.",
			InnerField: "id",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "prompt.variables",
			InnerField: "variables",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "prompt.version",
			Usage:      "Optional version identifier of the stored prompt.",
			InnerField: "version",
		},
	},
})

func handleResponsesCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomdedaluslabsdedalussdkgo.ResponseNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Responses.New(ctx, params, options...)
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
		Title:          "responses create",
		Transform:      transform,
	})
}
