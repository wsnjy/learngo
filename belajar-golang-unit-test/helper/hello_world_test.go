package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("west")
	}
}

func BenchmarkTestSub(b *testing.B) {

	b.Run("Bobbi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bobbi")
		}
	})

	b.Run("Bikul", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Bikul")
		}
	})
}

func BenchmarkTableHelloWorld(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Bobi",
			request: "Bobi",
		},
		{
			name:    "Bikul",
			request: "Bikul",
		},
		{
			name:    "Jering",
			request: "Jering",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("west")
	assert.Equal(t, result, "Hello west", "result should be Hello west")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("west")
	require.Equal(t, result, "Hello west", "result should be Hello west")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Test ini tidak dapat jalan pada OS MAC")
	}

	result := HelloWorld("west")
	require.Equal(t, result, "Hello west", "result should be Hello west")
}

func TestMain(m *testing.M) {
	fmt.Println("Sebelum test")
	m.Run()
	fmt.Println("Setelah test")
}

func TestSubTest(t *testing.T) {

	t.Run("Bobbi", func(t *testing.T) {
		result := HelloWorld("Bobbi")
		require.Equal(t, result, "Hello Bobbi", "result should be Hello Bobbi")
	})

	t.Run("Bikul", func(t *testing.T) {
		result := HelloWorld("Bikul")
		require.Equal(t, result, "Hello Bikul", "result should be Hello Bikul")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Bobi",
			request:  "Bobi",
			expected: "Hello Bobi",
		},
		{
			name:     "Bikul",
			request:  "Bikul",
			expected: "Hello Bikul",
		},
		{
			name:     "Jering",
			request:  "Jering",
			expected: "Hello Jering",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result, "result should be", test.expected)
		})
	}
}
