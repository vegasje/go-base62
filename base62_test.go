package base62

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		length         int
		expectedLength int
	}{
		{-1, 0},
		{0, 0},
		{1, 1},
		{100, 100},
	}
	for _, test := range tests {
		actual := Generate(test.length)
		require.Equal(t, test.expectedLength, len(actual), "Generated string is of the incorrect length.")
	}
}
