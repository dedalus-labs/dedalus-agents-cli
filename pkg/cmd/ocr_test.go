// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dedalus-labs/dedalus-agents-cli/internal/mocktest"
	"github.com/dedalus-labs/dedalus-agents-cli/internal/requestflag"
)

func TestOCRProcess(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ocr", "process",
			"--document", "{document_url: document_url, type: type}",
			"--model", "model",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(ocrProcess)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"ocr", "process",
			"--document.document-url", "document_url",
			"--document.type", "type",
			"--model", "model",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"document:\n" +
			"  document_url: document_url\n" +
			"  type: type\n" +
			"model: model\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"ocr", "process",
		)
	})
}
