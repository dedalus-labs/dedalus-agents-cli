// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/dedalus-sdk-cli/internal/mocktest"
	"github.com/stainless-sdks/dedalus-sdk-cli/internal/requestflag"
)

func TestResponsesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"responses", "create",
			"--background=true",
			"--conversation", "string",
			"--credentials", "{connection_name: external-service, values: {api_key: sk-...}}",
			"--frequency-penalty", "0",
			"--include", "[message.output_text.logprobs]",
			"--input", "What is the capital of France?",
			"--instructions", "You are a helpful assistant.",
			"--max-output-tokens", "1000",
			"--max-tool-calls", "10",
			"--mcp-servers", "dedalus-labs/example-server",
			"--metadata", "{foo: string}",
			"--model", "openai/gpt-4o",
			"--parallel-tool-calls=true",
			"--presence-penalty", "0",
			"--previous-response-id", "previous_response_id",
			"--prompt", "{id: id, variables: {foo: string}, version: version}",
			"--prompt-cache-key", "prompt_cache_key",
			"--reasoning", "{foo: string}",
			"--safety-identifier", "safety_identifier",
			"--service-tier", "auto",
			"--store=true",
			"--stream=true",
			"--stream-options", "{include_usage: true}",
			"--temperature", "0",
			"--text", "{type: text}",
			"--tool-choice", "auto",
			"--tool", "[{function: {description: null, name: null, parameters: null}, type: function}]",
			"--top-logprobs", "5",
			"--top-p", "0.1",
			"--truncation", "auto",
			"--user", "user",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(responsesCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"responses", "create",
			"--background=true",
			"--conversation", "string",
			"--credentials", "{connection_name: external-service, values: {api_key: sk-...}}",
			"--frequency-penalty", "0",
			"--include", "[message.output_text.logprobs]",
			"--input", "What is the capital of France?",
			"--instructions", "You are a helpful assistant.",
			"--max-output-tokens", "1000",
			"--max-tool-calls", "10",
			"--mcp-servers", "dedalus-labs/example-server",
			"--metadata", "{foo: string}",
			"--model", "openai/gpt-4o",
			"--parallel-tool-calls=true",
			"--presence-penalty", "0",
			"--previous-response-id", "previous_response_id",
			"--prompt.id", "id",
			"--prompt.variables", "{foo: string}",
			"--prompt.version", "version",
			"--prompt-cache-key", "prompt_cache_key",
			"--reasoning", "{foo: string}",
			"--safety-identifier", "safety_identifier",
			"--service-tier", "auto",
			"--store=true",
			"--stream=true",
			"--stream-options", "{include_usage: true}",
			"--temperature", "0",
			"--text", "{type: text}",
			"--tool-choice", "auto",
			"--tool", "[{function: {description: null, name: null, parameters: null}, type: function}]",
			"--top-logprobs", "5",
			"--top-p", "0.1",
			"--truncation", "auto",
			"--user", "user",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"background: true\n" +
			"conversation: string\n" +
			"credentials:\n" +
			"  connection_name: external-service\n" +
			"  values:\n" +
			"    api_key: sk-...\n" +
			"frequency_penalty: 0\n" +
			"include:\n" +
			"  - message.output_text.logprobs\n" +
			"input: What is the capital of France?\n" +
			"instructions: You are a helpful assistant.\n" +
			"max_output_tokens: 1000\n" +
			"max_tool_calls: 10\n" +
			"mcp_servers: dedalus-labs/example-server\n" +
			"metadata:\n" +
			"  foo: string\n" +
			"model: openai/gpt-4o\n" +
			"parallel_tool_calls: true\n" +
			"presence_penalty: 0\n" +
			"previous_response_id: previous_response_id\n" +
			"prompt:\n" +
			"  id: id\n" +
			"  variables:\n" +
			"    foo: string\n" +
			"  version: version\n" +
			"prompt_cache_key: prompt_cache_key\n" +
			"reasoning:\n" +
			"  foo: string\n" +
			"safety_identifier: safety_identifier\n" +
			"service_tier: auto\n" +
			"store: true\n" +
			"stream: true\n" +
			"stream_options:\n" +
			"  include_usage: true\n" +
			"temperature: 0\n" +
			"text:\n" +
			"  type: text\n" +
			"tool_choice: auto\n" +
			"tools:\n" +
			"  - function:\n" +
			"      description: null\n" +
			"      name: null\n" +
			"      parameters: null\n" +
			"    type: function\n" +
			"top_logprobs: 5\n" +
			"top_p: 0.1\n" +
			"truncation: auto\n" +
			"user: user\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"responses", "create",
		)
	})
}
