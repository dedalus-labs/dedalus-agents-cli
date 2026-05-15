// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dedalus-labs/dedalus-agents-cli/internal/mocktest"
)

func TestEmbeddingsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"embeddings", "create",
			"--input", "string",
			"--model", "text-embedding-ada-002",
			"--dimensions", "1",
			"--encoding-format", "float",
			"--user", "user",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"input: string\n" +
			"model: text-embedding-ada-002\n" +
			"dimensions: 1\n" +
			"encoding_format: float\n" +
			"user: user\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"embeddings", "create",
		)
	})
}
