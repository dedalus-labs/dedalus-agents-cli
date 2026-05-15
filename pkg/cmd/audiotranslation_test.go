// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/dedalus-labs/dedalus-agents-cli/internal/mocktest"
)

func TestAudioTranslationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"audio:translations", "create",
			"--file", mocktest.TestFile(t, "Example data"),
			"--model", "model",
			"--prompt", "prompt",
			"--response-format", "response_format",
			"--temperature", "0",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"file: Example data\n" +
			"model: model\n" +
			"prompt: prompt\n" +
			"response_format: response_format\n" +
			"temperature: 0\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"audio:translations", "create",
		)
	})
}
