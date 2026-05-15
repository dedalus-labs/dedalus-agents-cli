// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/dedalus-sdk-cli/internal/mocktest"
	"github.com/stainless-sdks/dedalus-sdk-cli/internal/requestflag"
)

func TestChatCompletionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"chat:completions", "create",
			"--max-items", "10",
			"--model", "openai/gpt-5",
			"--agent-attributes", "{accuracy: 0.9, complexity: 0.8}",
			"--audio", "{format: mp3, voice: alloy}",
			"--automatic-tool-execution=true",
			"--cached-content", "cached_content",
			"--correlation-id", "correlation_id",
			"--credentials", "{connection_name: external-service, values: {api_key: sk-...}}",
			"--deferred=true",
			"--deferred-call", "[{id: id, name: name, arguments: {foo: string}, blocked_by: [string], dependencies: [string], venue: venue}]",
			"--frequency-penalty", "-2",
			"--function-call", "function_call",
			"--function", "[{name: name, description: description, parameters: {foo: bar}}]",
			"--generation-config", "{foo: string}",
			"--guardrail", "[{foo: bar}]",
			"--handoff-config", "{foo: bar}",
			"--handoff-mode=true",
			"--inference-geo", "inference_geo",
			"--logit-bias", "{foo: 0}",
			"--logprobs=true",
			"--max-completion-tokens", "0",
			"--max-tokens", "1",
			"--max-turns", "5",
			"--mcp-servers", "dedalus-labs/example-server",
			"--message", "[{content: string, role: developer, name: name}]",
			"--metadata", "{foo: string}",
			"--modality", "[string]",
			"--model-attributes", "{gpt-5: {accuracy: 0.95, speed: 0.6}}",
			"--n", "1",
			"--output-config", "{foo: string}",
			"--parallel-tool-calls=true",
			"--prediction", "{content: string, type: content}",
			"--presence-penalty", "-2",
			"--prompt-cache-key", "prompt_cache_key",
			"--prompt-cache-retention", "prompt_cache_retention",
			"--prompt-mode", "reasoning",
			"--reasoning-effort", "reasoning_effort",
			"--response-format", "{type: text}",
			"--safe-prompt=true",
			"--safety-identifier", "safety_identifier",
			"--safety-setting", "[{category: HARM_CATEGORY_UNSPECIFIED, threshold: HARM_BLOCK_THRESHOLD_UNSPECIFIED}]",
			"--search-parameters", "{foo: string}",
			"--seed", "0",
			"--service-tier", "service_tier",
			"--speed", "standard",
			"--stop", "[string]",
			"--store=true",
			"--stream=false",
			"--stream-options", "{foo: string}",
			"--system-instruction", "{foo: string}",
			"--temperature", "0",
			"--thinking", "{budget_tokens: 1024, type: enabled}",
			"--tool-choice", "string",
			"--tool-config", "{foo: string}",
			"--tool", "[{function: {name: name}, type: function}]",
			"--top-k", "0",
			"--top-logprobs", "0",
			"--top-p", "0",
			"--user", "user",
			"--verbosity", "verbosity",
			"--web-search-options", "{foo: string}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(chatCompletionsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"chat:completions", "create",
			"--max-items", "10",
			"--model", "openai/gpt-5",
			"--agent-attributes", "{accuracy: 0.9, complexity: 0.8}",
			"--audio.format", "mp3",
			"--audio.voice", "alloy",
			"--automatic-tool-execution=true",
			"--cached-content", "cached_content",
			"--correlation-id", "correlation_id",
			"--credentials", "{connection_name: external-service, values: {api_key: sk-...}}",
			"--deferred=true",
			"--deferred-call.id", "id",
			"--deferred-call.name", "name",
			"--deferred-call.arguments", "{foo: string}",
			"--deferred-call.blocked-by", "[string]",
			"--deferred-call.dependencies", "[string]",
			"--deferred-call.venue", "venue",
			"--frequency-penalty", "-2",
			"--function-call", "function_call",
			"--function.name", "name",
			"--function.description", "description",
			"--function.parameters", "{foo: bar}",
			"--generation-config", "{foo: string}",
			"--guardrail", "[{foo: bar}]",
			"--handoff-config", "{foo: bar}",
			"--handoff-mode=true",
			"--inference-geo", "inference_geo",
			"--logit-bias", "{foo: 0}",
			"--logprobs=true",
			"--max-completion-tokens", "0",
			"--max-tokens", "1",
			"--max-turns", "5",
			"--mcp-servers", "dedalus-labs/example-server",
			"--message", "[{content: string, role: developer, name: name}]",
			"--metadata", "{foo: string}",
			"--modality", "[string]",
			"--model-attributes", "{gpt-5: {accuracy: 0.95, speed: 0.6}}",
			"--n", "1",
			"--output-config", "{foo: string}",
			"--parallel-tool-calls=true",
			"--prediction.content", "string",
			"--prediction.type", "content",
			"--presence-penalty", "-2",
			"--prompt-cache-key", "prompt_cache_key",
			"--prompt-cache-retention", "prompt_cache_retention",
			"--prompt-mode", "reasoning",
			"--reasoning-effort", "reasoning_effort",
			"--response-format", "{type: text}",
			"--safe-prompt=true",
			"--safety-identifier", "safety_identifier",
			"--safety-setting.category", "HARM_CATEGORY_UNSPECIFIED",
			"--safety-setting.threshold", "HARM_BLOCK_THRESHOLD_UNSPECIFIED",
			"--search-parameters", "{foo: string}",
			"--seed", "0",
			"--service-tier", "service_tier",
			"--speed", "standard",
			"--stop", "[string]",
			"--store=true",
			"--stream=false",
			"--stream-options", "{foo: string}",
			"--system-instruction", "{foo: string}",
			"--temperature", "0",
			"--thinking", "{budget_tokens: 1024, type: enabled}",
			"--tool-choice", "string",
			"--tool-config", "{foo: string}",
			"--tool.function", "{name: name}",
			"--tool.type", "function",
			"--top-k", "0",
			"--top-logprobs", "0",
			"--top-p", "0",
			"--user", "user",
			"--verbosity", "verbosity",
			"--web-search-options", "{foo: string}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"model: openai/gpt-5\n" +
			"agent_attributes:\n" +
			"  accuracy: 0.9\n" +
			"  complexity: 0.8\n" +
			"audio:\n" +
			"  format: mp3\n" +
			"  voice: alloy\n" +
			"automatic_tool_execution: true\n" +
			"cached_content: cached_content\n" +
			"correlation_id: correlation_id\n" +
			"credentials:\n" +
			"  connection_name: external-service\n" +
			"  values:\n" +
			"    api_key: sk-...\n" +
			"deferred: true\n" +
			"deferred_calls:\n" +
			"  - id: id\n" +
			"    name: name\n" +
			"    arguments:\n" +
			"      foo: string\n" +
			"    blocked_by:\n" +
			"      - string\n" +
			"    dependencies:\n" +
			"      - string\n" +
			"    venue: venue\n" +
			"frequency_penalty: -2\n" +
			"function_call: function_call\n" +
			"functions:\n" +
			"  - name: name\n" +
			"    description: description\n" +
			"    parameters:\n" +
			"      foo: bar\n" +
			"generation_config:\n" +
			"  foo: string\n" +
			"guardrails:\n" +
			"  - foo: bar\n" +
			"handoff_config:\n" +
			"  foo: bar\n" +
			"handoff_mode: true\n" +
			"inference_geo: inference_geo\n" +
			"logit_bias:\n" +
			"  foo: 0\n" +
			"logprobs: true\n" +
			"max_completion_tokens: 0\n" +
			"max_tokens: 1\n" +
			"max_turns: 5\n" +
			"mcp_servers: dedalus-labs/example-server\n" +
			"messages:\n" +
			"  - content: string\n" +
			"    role: developer\n" +
			"    name: name\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"modalities:\n" +
			"  - string\n" +
			"model_attributes:\n" +
			"  gpt-5:\n" +
			"    accuracy: 0.95\n" +
			"    speed: 0.6\n" +
			"'n': 1\n" +
			"output_config:\n" +
			"  foo: string\n" +
			"parallel_tool_calls: true\n" +
			"prediction:\n" +
			"  content: string\n" +
			"  type: content\n" +
			"presence_penalty: -2\n" +
			"prompt_cache_key: prompt_cache_key\n" +
			"prompt_cache_retention: prompt_cache_retention\n" +
			"prompt_mode: reasoning\n" +
			"reasoning_effort: reasoning_effort\n" +
			"response_format:\n" +
			"  type: text\n" +
			"safe_prompt: true\n" +
			"safety_identifier: safety_identifier\n" +
			"safety_settings:\n" +
			"  - category: HARM_CATEGORY_UNSPECIFIED\n" +
			"    threshold: HARM_BLOCK_THRESHOLD_UNSPECIFIED\n" +
			"search_parameters:\n" +
			"  foo: string\n" +
			"seed: 0\n" +
			"service_tier: service_tier\n" +
			"speed: standard\n" +
			"stop:\n" +
			"  - string\n" +
			"store: true\n" +
			"stream: false\n" +
			"stream_options:\n" +
			"  foo: string\n" +
			"system_instruction:\n" +
			"  foo: string\n" +
			"temperature: 0\n" +
			"thinking:\n" +
			"  budget_tokens: 1024\n" +
			"  type: enabled\n" +
			"tool_choice: string\n" +
			"tool_config:\n" +
			"  foo: string\n" +
			"tools:\n" +
			"  - function:\n" +
			"      name: name\n" +
			"    type: function\n" +
			"top_k: 0\n" +
			"top_logprobs: 0\n" +
			"top_p: 0\n" +
			"user: user\n" +
			"verbosity: verbosity\n" +
			"web_search_options:\n" +
			"  foo: string\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"chat:completions", "create",
			"--max-items", "10",
		)
	})
}
