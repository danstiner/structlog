package messagetemplates

import (
	"fmt"
	"testing"

	structlog "github.com/danstiner/go-structlog"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ExampleFormat() {
	pos := struct {
		Lat  float32
		Long float32
	}{
		Lat:  25,
		Long: 132,
	}
	elapsed := 34
	msg, _, err := Format("Processed {@position} in {elapsed} ms", pos, elapsed)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(msg)
}

func TestFormatString(t *testing.T) {
	msg, m, err := Format("Message")
	require.NoError(t, err)
	assert.Equal(t, "Message", msg)
	assert.Equal(t, []structlog.KV{}, m)
}

func TestFormatHole(t *testing.T) {
	msg, m, err := Format("{msg}", "Hello world!")
	require.NoError(t, err)
	assert.Equal(t, "Hello world!", msg)
	assert.Equal(t, []structlog.KV{
		{"msg", "Hello world!"},
	}, m)
}

func TestFormatSerializeHole(t *testing.T) {
	pos := struct {
		Lat  float32
		Long float32
	}{
		Lat:  25,
		Long: 132,
	}
	msg, m, err := Format("Processed {@position}", pos)
	require.NoError(t, err)
	assert.Equal(t, `Processed {"Lat":25,"Long":132}`, msg)
	assert.Equal(t, []structlog.KV{
		{"position", pos},
	}, m)
}

func TestFormatTooFewArgs(t *testing.T) {
	_, _, err := Format("{msg}")
	require.Error(t, err)
}

func TestFormatTooManyArgs(t *testing.T) {
	_, _, err := Format("{msg}", "Hello", "world!")
	require.Error(t, err)
}

// always store the result to a package level variable
// so the compiler cannot eliminate the Benchmark itself.
var result string
var resultMap []structlog.KV

func BenchmarkSprintfString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = fmt.Sprintf("Message")
	}
}

func BenchmarkFormatString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result, _, _ = Format("Message")
	}
}

func BenchmarkSprintfHole(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result = fmt.Sprintf("%v", "Hello world!")
	}
}

func BenchmarkFormatHole(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result, _, _ = Format("{msg}", "Hello world!")
	}
}

func BenchmarkRenderHole(b *testing.B) {
	template, err := Parse("{msg}")
	require.NoError(b, err)

	for n := 0; n < b.N; n++ {
		result, _, _ = Render(template, "Hello world!")
	}
}

func BenchmarkFormatSerializeHole(b *testing.B) {
	pos := struct {
		Lat  float32
		Long float32
	}{
		Lat:  25,
		Long: 132,
	}

	for n := 0; n < b.N; n++ {
		result, _, _ = Format("Processed {@position} in 34 ms", pos)
	}
}

func BenchmarkFormat10Fields(b *testing.B) {
	for n := 0; n < b.N; n++ {
		result, _, _ = Format("Ten fields {0} {1} {2} {3} {4} {5} {6} {7} {8} {9}", 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}

func BenchmarkRender10Fields(b *testing.B) {
	template, err := Parse("Ten fields {0} {1} {2} {3} {4} {5} {6} {7} {8} {9}")
	require.NoError(b, err)

	for n := 0; n < b.N; n++ {
		result, _, _ = Render(template, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	}
}
