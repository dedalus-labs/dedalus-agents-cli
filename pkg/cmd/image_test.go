// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"strings"
	"testing"

	"github.com/stainless-sdks/dedalus-sdk-cli/internal/mocktest"
)

func TestImagesCreateVariation(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"images", "create-variation",
			"--image", mocktest.TestFile(t, "Example data"),
			"--model", "model",
			"--n", "0",
			"--response-format", "response_format",
			"--size", "size",
			"--user", "user",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"image: Example data\n" +
			"model: model\n" +
			"'n': 0\n" +
			"response_format: response_format\n" +
			"size: size\n" +
			"user: user\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"images", "create-variation",
		)
	})
}

func TestImagesEdit(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"images", "edit",
			"--image", mocktest.TestFile(t, "Example data"),
			"--prompt", "prompt",
			"--mask", mocktest.TestFile(t, "Example data"),
			"--model", "model",
			"--n", "0",
			"--response-format", "response_format",
			"--size", "size",
			"--user", "user",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		testFile := mocktest.TestFile(t, "Example data")
		// Test piping YAML data over stdin
		pipeDataStr := "" +
			"image: Example data\n" +
			"prompt: prompt\n" +
			"mask: Example data\n" +
			"model: model\n" +
			"'n': 0\n" +
			"response_format: response_format\n" +
			"size: size\n" +
			"user: user\n"
		pipeDataStr = strings.ReplaceAll(pipeDataStr, "Example data", testFile)
		pipeData := []byte(pipeDataStr)
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"images", "edit",
		)
	})
}

func TestImagesGenerate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"images", "generate",
			"--prompt", "A white siamese cat",
			"--background", "transparent",
			"--model", "openai/dall-e-3",
			"--moderation", "auto",
			"--n", "1",
			"--output-compression", "85",
			"--output-format", "png",
			"--partial-images", "0",
			"--quality", "standard",
			"--response-format", "url",
			"--size", "1024x1024",
			"--stream=true",
			"--style", "vivid",
			"--user", "user",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"prompt: A white siamese cat\n" +
			"background: transparent\n" +
			"model: openai/dall-e-3\n" +
			"moderation: auto\n" +
			"'n': 1\n" +
			"output_compression: 85\n" +
			"output_format: png\n" +
			"partial_images: 0\n" +
			"quality: standard\n" +
			"response_format: url\n" +
			"size: 1024x1024\n" +
			"stream: true\n" +
			"style: vivid\n" +
			"user: user\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"images", "generate",
		)
	})
}
