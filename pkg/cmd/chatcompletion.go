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

var chatCompletionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a chat completion.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "model",
			Usage:    "Model identifier. Accepts model ID strings, lists for routing, or DedalusModel objects with per-model settings.",
			Required: true,
			BodyPath: "model",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "agent-attributes",
			Usage:    "Agent attributes. Values in [0.0, 1.0].",
			BodyPath: "agent_attributes",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "audio",
			Usage:    "Parameters for audio output. Required when audio output is requested with\n`modalities: [\"audio\"]`. [Learn more](/docs/guides/audio).\n\nFields:\n- voice (required): VoiceIdsOrCustomVoice\n- format (required): Literal[\"wav\", \"aac\", \"mp3\", \"flac\", \"opus\", \"pcm16\"]",
			BodyPath: "audio",
		},
		&requestflag.Flag[bool]{
			Name:     "automatic-tool-execution",
			Usage:    "Execute tools server-side. If false, returns raw tool calls for manual handling.",
			Default:  true,
			BodyPath: "automatic_tool_execution",
		},
		&requestflag.Flag[*string]{
			Name:     "cached-content",
			Usage:    "Optional. The name of the content [cached](https://ai.google.dev/gemini-api/docs/caching) to use as context to serve the prediction. Format: `cachedContents/{cachedContent}`",
			BodyPath: "cached_content",
		},
		&requestflag.Flag[*string]{
			Name:     "correlation-id",
			Usage:    "Stable session ID for resuming a previous handoff. Returned by the server on handoff; echo it on the next request to resume.",
			BodyPath: "correlation_id",
		},
		&requestflag.Flag[any]{
			Name:     "credentials",
			Usage:    "Credentials for MCP server authentication. Each credential is matched to servers by connection name.",
			BodyPath: "credentials",
		},
		&requestflag.Flag[*bool]{
			Name:     "deferred",
			Usage:    "If set to `true`, the request returns a `request_id`. You can then get the deferred response by GET `/v1/chat/deferred-completion/{request_id}`.",
			Default:  requestflag.Ptr[bool](false),
			BodyPath: "deferred",
		},
		&requestflag.Flag[any]{
			Name:     "deferred-call",
			Usage:    "Tier 2 stateless resumption. Deferred tool specs from a previous handoff response, sent back verbatim so the server can resume without Redis.",
			BodyPath: "deferred_calls",
		},
		&requestflag.Flag[*float64]{
			Name:     "frequency-penalty",
			Usage:    "Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim. ",
			Default:  requestflag.Ptr[float64](0),
			BodyPath: "frequency_penalty",
		},
		&requestflag.Flag[*string]{
			Name:     "function-call",
			Usage:    "Deprecated in favor of `tool_choice`.  Controls which (if any) function is called by the model.  `none` means the model will not call a function and instead generates a message.  `auto` means the model can pick between generating a message or calling a function.  Specifying a particular function via `{\"name\": \"my_function\"}` forces the model to call that function.  `none` is the default when no functions are present. `auto` is the default if functions are present. ",
			BodyPath: "function_call",
		},
		&requestflag.Flag[any]{
			Name:     "function",
			Usage:    "Deprecated in favor of `tools`.  A list of functions the model may generate JSON inputs for. ",
			BodyPath: "functions",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "generation-config",
			BodyPath: "generation_config",
		},
		&requestflag.Flag[any]{
			Name:     "guardrail",
			Usage:    "Content filtering and safety policy configuration.",
			BodyPath: "guardrails",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "handoff-config",
			Usage:    "Configuration for multi-model handoffs.",
			BodyPath: "handoff_config",
		},
		&requestflag.Flag[*bool]{
			Name:     "handoff-mode",
			Usage:    "Handoff control. None or omitted: auto-detect. true: structured handoff (SDK). false: drop-in (LLM re-run for mixed turns).",
			BodyPath: "handoff_mode",
		},
		&requestflag.Flag[*string]{
			Name:     "inference-geo",
			Usage:    "Specifies the geographic region for inference processing. If not specified, the workspace's `default_inference_geo` is used.",
			BodyPath: "inference_geo",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "logit-bias",
			Usage:    "Modify the likelihood of specified tokens appearing in the completion.  Accepts a JSON object that maps tokens (specified by their token ID in the tokenizer) to an associated bias value from -100 to 100. Mathematically, the bias is added to the logits generated by the model prior to sampling. The exact effect will vary per model, but values between -1 and 1 should decrease or increase likelihood of selection; values like -100 or 100 should result in a ban or exclusive selection of the relevant token. ",
			BodyPath: "logit_bias",
		},
		&requestflag.Flag[*bool]{
			Name:     "logprobs",
			Usage:    "Whether to return log probabilities of the output tokens or not. If true, returns the log probabilities of each output token returned in the `content` of `message`. ",
			Default:  requestflag.Ptr[bool](false),
			BodyPath: "logprobs",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-completion-tokens",
			Usage:    "Maximum tokens in completion (newer parameter name)",
			BodyPath: "max_completion_tokens",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-tokens",
			Usage:    "Maximum tokens in completion",
			BodyPath: "max_tokens",
		},
		&requestflag.Flag[*int64]{
			Name:     "max-turns",
			Usage:    "Maximum conversation turns.",
			BodyPath: "max_turns",
		},
		&requestflag.Flag[any]{
			Name:     "mcp-servers",
			Usage:    "MCP server identifiers. Accepts marketplace slugs, URLs, or MCPServerSpec objects. MCP tools are executed server-side and billed separately.",
			BodyPath: "mcp_servers",
		},
		&requestflag.Flag[any]{
			Name:     "message",
			Usage:    "Conversation history (OpenAI: messages, Google: contents, Responses: input)",
			BodyPath: "messages",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[any]{
			Name:     "modality",
			Usage:    "Output types that you would like the model to generate. Most models are capable of generating text, which is the default:  `[\"text\"]`  The `gpt-4o-audio-preview` model can also be used to [generate audio](/docs/guides/audio). To request that this model generate both text and audio responses, you can use:  `[\"text\", \"audio\"]` ",
			BodyPath: "modalities",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "model-attributes",
			Usage:    "Model attributes for routing. Maps model IDs to attribute dictionaries with values in [0.0, 1.0].",
			BodyPath: "model_attributes",
		},
		&requestflag.Flag[*int64]{
			Name:     "n",
			Usage:    "How many chat completion choices to generate for each input message. Note that you will be charged based on the number of generated tokens across all of the choices. Keep `n` as `1` to minimize costs.",
			Default:  requestflag.Ptr[int64](1),
			BodyPath: "n",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "output-config",
			BodyPath: "output_config",
		},
		&requestflag.Flag[*bool]{
			Name:     "parallel-tool-calls",
			Usage:    "Whether to enable parallel tool calls (Anthropic uses inverted polarity).",
			Default:  requestflag.Ptr[bool](true),
			BodyPath: "parallel_tool_calls",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "prediction",
			Usage:    "Static predicted output content, such as the content of a text file that is\nbeing regenerated.\n\nFields:\n- type (required): Literal[\"content\"]\n- content (required): str | Annotated[list[ChatCompletionRequestMessageContentPartText], MinLen(1), ArrayTitle(\"PredictionContentArray\")]",
			BodyPath: "prediction",
		},
		&requestflag.Flag[*float64]{
			Name:     "presence-penalty",
			Usage:    "Number between -2.0 and 2.0. Positive values penalize new tokens based on whether they appear in the text so far, increasing the model's likelihood to talk about new topics. ",
			Default:  requestflag.Ptr[float64](0),
			BodyPath: "presence_penalty",
		},
		&requestflag.Flag[*string]{
			Name:     "prompt-cache-key",
			Usage:    "Used by OpenAI to cache responses for similar requests to optimize your cache hit rates. Replaces the `user` field. [Learn more](/docs/guides/prompt-caching). ",
			BodyPath: "prompt_cache_key",
		},
		&requestflag.Flag[*string]{
			Name:     "prompt-cache-retention",
			Usage:    "The retention policy for the prompt cache. Set to `24h` to enable extended prompt caching, which keeps cached prefixes active for longer, up to a maximum of 24 hours. [Learn more](/docs/guides/prompt-caching#prompt-cache-retention). ",
			BodyPath: "prompt_cache_retention",
		},
		&requestflag.Flag[*string]{
			Name:     "prompt-mode",
			Usage:    "Allows toggling between the reasoning mode and no system prompt. When set to `reasoning` the system prompt for reasoning models will be used.",
			BodyPath: "prompt_mode",
		},
		&requestflag.Flag[*string]{
			Name:     "reasoning-effort",
			Usage:    "Constrains effort on reasoning for [reasoning models](https://platform.openai.com/docs/guides/reasoning). Currently supported values are `none`, `minimal`, `low`, `medium`, `high`, and `xhigh`. Reducing reasoning effort can result in faster responses and fewer tokens used on reasoning in a response.  - `gpt-5.1` defaults to `none`, which does not perform reasoning. The supported reasoning values for `gpt-5.1` are `none`, `low`, `medium`, and `high`. Tool calls are supported for all reasoning values in gpt-5.1. - All models before `gpt-5.1` default to `medium` reasoning effort, and do not support `none`. - The `gpt-5-pro` model defaults to (and only supports) `high` reasoning effort. - `xhigh` is supported for all models after `gpt-5.1-codex-max`. ",
			Default:  requestflag.Ptr[string]("medium"),
			BodyPath: "reasoning_effort",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "response-format",
			Usage:    "An object specifying the format that the model must output.  Setting to `{ \"type\": \"json_schema\", \"json_schema\": {...} }` enables Structured Outputs which ensures the model will match your supplied JSON schema. Learn more in the [Structured Outputs guide](/docs/guides/structured-outputs).  Setting to `{ \"type\": \"json_object\" }` enables the older JSON mode, which ensures the message the model generates is valid JSON. Using `json_schema` is preferred for models that support it. ",
			BodyPath: "response_format",
		},
		&requestflag.Flag[*bool]{
			Name:     "safe-prompt",
			Usage:    "Whether to inject a safety prompt before all conversations.",
			Default:  requestflag.Ptr[bool](false),
			BodyPath: "safe_prompt",
		},
		&requestflag.Flag[*string]{
			Name:     "safety-identifier",
			Usage:    "A stable identifier used to help detect users of your application that may be violating OpenAI's usage policies. The IDs should be a string that uniquely identifies each user. We recommend hashing their username or email address, in order to avoid sending us any identifying information. [Learn more](/docs/guides/safety-best-practices#safety-identifiers). ",
			BodyPath: "safety_identifier",
		},
		&requestflag.Flag[any]{
			Name:     "safety-setting",
			Usage:    "Safety/content filtering settings (Google-specific)",
			BodyPath: "safety_settings",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "search-parameters",
			BodyPath: "search_parameters",
		},
		&requestflag.Flag[*int64]{
			Name:     "seed",
			Usage:    "Random seed for deterministic output",
			BodyPath: "seed",
		},
		&requestflag.Flag[*string]{
			Name:     "service-tier",
			Usage:    "Service tier for request processing",
			Default:  requestflag.Ptr[string]("auto"),
			BodyPath: "service_tier",
		},
		&requestflag.Flag[*string]{
			Name:     "speed",
			Usage:    "The inference speed mode for this request. `\"fast\"` enables high output-tokens-per-second inference.",
			BodyPath: "speed",
		},
		&requestflag.Flag[any]{
			Name:     "stop",
			Usage:    "Sequences that stop generation",
			BodyPath: "stop",
		},
		&requestflag.Flag[*bool]{
			Name:     "store",
			Usage:    "Whether or not to store the output of this chat completion request for use in our [model distillation](/docs/guides/distillation) or [evals](/docs/guides/evals) products.  Supports text and image inputs. Note: image inputs over 8MB will be dropped. ",
			Default:  requestflag.Ptr[bool](false),
			BodyPath: "store",
		},
		&requestflag.Flag[*bool]{
			Name:     "stream",
			Usage:    "Enable streaming response",
			Default:  requestflag.Ptr[bool](false),
			BodyPath: "stream",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "stream-options",
			BodyPath: "stream_options",
		},
		&requestflag.Flag[any]{
			Name:     "system-instruction",
			Usage:    "System instruction/prompt",
			BodyPath: "system_instruction",
		},
		&requestflag.Flag[*float64]{
			Name:     "temperature",
			Usage:    "Sampling temperature (0-2 for most providers)",
			Default:  requestflag.Ptr[float64](1),
			BodyPath: "temperature",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "thinking",
			Usage:    "Extended thinking configuration (Anthropic-specific)",
			BodyPath: "thinking",
		},
		&requestflag.Flag[any]{
			Name:     "tool-choice",
			Usage:    "Controls which (if any) tool is called by the model. `none` means the model will not call any tool and instead generates a message. `auto` means the model can pick between generating a message or calling one or more tools. `required` means the model must call one or more tools. Specifying a particular tool via `{\"type\": \"function\", \"function\": {\"name\": \"my_function\"}}` forces the model to call that tool.  `none` is the default when no tools are present. `auto` is the default if tools are present. ",
			Default:  "auto",
			BodyPath: "tool_choice",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "tool-config",
			BodyPath: "tool_config",
		},
		&requestflag.Flag[any]{
			Name:     "tool",
			Usage:    "Available tools/functions for the model",
			BodyPath: "tools",
		},
		&requestflag.Flag[*int64]{
			Name:     "top-k",
			Usage:    "Top-k sampling parameter",
			BodyPath: "top_k",
		},
		&requestflag.Flag[*int64]{
			Name:     "top-logprobs",
			Usage:    "An integer between 0 and 20 specifying the number of most likely tokens to return at each token position, each with an associated log probability. `logprobs` must be set to `true` if this parameter is used. ",
			BodyPath: "top_logprobs",
		},
		&requestflag.Flag[*float64]{
			Name:     "top-p",
			Usage:    "Nucleus sampling threshold",
			Default:  requestflag.Ptr[float64](1),
			BodyPath: "top_p",
		},
		&requestflag.Flag[*string]{
			Name:     "user",
			Usage:    "This field is being replaced by `safety_identifier` and `prompt_cache_key`. Use `prompt_cache_key` instead to maintain caching optimizations. A stable identifier for your end-users. Used to boost cache hit rates by better bucketing similar requests and  to help OpenAI detect and prevent abuse. [Learn more](/docs/guides/safety-best-practices#safety-identifiers). ",
			BodyPath: "user",
		},
		&requestflag.Flag[*string]{
			Name:     "verbosity",
			Usage:    "Constrains the verbosity of the model's response. Lower values will result in more concise responses, while higher values will result in more verbose responses. Currently supported values are `low`, `medium`, and `high`. ",
			Default:  requestflag.Ptr[string]("medium"),
			BodyPath: "verbosity",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "web-search-options",
			BodyPath: "web_search_options",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleChatCompletionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"audio": {
		&requestflag.InnerFlag[string]{
			Name:       "audio.format",
			Usage:      "Specifies the output audio format. Must be one of `wav`, `mp3`, `flac`,\n`opus`, or `pcm16`.",
			InnerField: "format",
		},
		&requestflag.InnerFlag[any]{
			Name:       "audio.voice",
			Usage:      "The voice the model uses to respond. Supported built-in voices are\n`alloy`, `ash`, `ballad`, `coral`, `echo`, `fable`, `nova`, `onyx`,\n`sage`, `shimmer`, `marin`, and `cedar`. You may also provide a\ncustom voice object with an `id`, for example `{ \"id\": \"voice_1234\" }`.",
			InnerField: "voice",
		},
	},
	"deferred-call": {
		&requestflag.InnerFlag[string]{
			Name:                  "deferred-call.id",
			Usage:                 "Unique identifier for this deferred call.",
			InnerField:            "id",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "deferred-call.name",
			Usage:                 "Name of the tool.",
			InnerField:            "name",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "deferred-call.arguments",
			InnerField:            "arguments",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[[]string]{
			Name:                  "deferred-call.blocked-by",
			Usage:                 "IDs of pending client calls blocking this call.",
			InnerField:            "blocked_by",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[[]string]{
			Name:                  "deferred-call.dependencies",
			Usage:                 "IDs of calls this depends on.",
			InnerField:            "dependencies",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "deferred-call.venue",
			Usage:                 "Execution venue (server or client).",
			InnerField:            "venue",
			OuterIsArrayOfObjects: true,
		},
	},
	"function": {
		&requestflag.InnerFlag[string]{
			Name:                  "function.name",
			Usage:                 "The name of the function to be called. Must be a-z, A-Z, 0-9, or contain underscores and dashes, with a maximum length of 64.",
			InnerField:            "name",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "function.description",
			Usage:                 "A description of what the function does, used by the model to choose when and how to call the function.",
			InnerField:            "description",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "function.parameters",
			Usage:                 "The parameters the functions accepts, described as a JSON Schema object. See the [guide](/docs/guides/function-calling) for examples, and the [JSON Schema reference](https://json-schema.org/understanding-json-schema/) for documentation about the format. \n\nOmitting `parameters` defines a function with an empty parameter list.",
			InnerField:            "parameters",
			OuterIsArrayOfObjects: true,
		},
	},
	"prediction": {
		&requestflag.InnerFlag[any]{
			Name:       "prediction.content",
			Usage:      "The content that should be matched when generating a model response.\nIf generated tokens would match this content, the entire model response\ncan be returned much more quickly.",
			InnerField: "content",
		},
		&requestflag.InnerFlag[string]{
			Name:       "prediction.type",
			Usage:      "The type of the predicted content you want to provide. This type is\ncurrently always `content`.",
			InnerField: "type",
		},
	},
	"safety-setting": {
		&requestflag.InnerFlag[string]{
			Name:                  "safety-setting.category",
			Usage:                 "Required. The category for this setting.",
			InnerField:            "category",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "safety-setting.threshold",
			Usage:                 "Required. Controls the probability threshold at which harm is blocked.",
			InnerField:            "threshold",
			OuterIsArrayOfObjects: true,
		},
	},
	"tool": {
		&requestflag.InnerFlag[map[string]any]{
			Name:                  "tool.function",
			Usage:                 "Schema for Function.\n\nFields:\n- name (required): str",
			InnerField:            "function",
			OuterIsArrayOfObjects: true,
		},
		&requestflag.InnerFlag[string]{
			Name:                  "tool.type",
			Usage:                 `Allowed values: "function".`,
			InnerField:            "type",
			OuterIsArrayOfObjects: true,
		},
	},
})

func handleChatCompletionsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := githubcomdedaluslabsdedalussdkgo.ChatCompletionNewParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if cmd.Bool("stream") {
		stream := client.Chat.Completions.NewStreaming(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(stream, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "chat:completions create",
			Transform:      transform,
		})
	} else {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Chat.Completions.New(ctx, params, options...)
		if err != nil {
			return err
		}

		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "chat:completions create",
			Transform:      transform,
		})
	}
}
