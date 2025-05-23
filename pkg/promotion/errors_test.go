package promotion

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsTerminal(t *testing.T) {
	testCases := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "not a terminal error",
			err:      errors.New("something went wrong"),
			expected: false,
		},
		{
			name:     "a terminal error",
			err:      &TerminalError{Err: errors.New("something went wrong")},
			expected: true,
		},
		{
			name: "a wrapped terminal error",
			err: fmt.Errorf(
				"an error occurred: %w",
				&TerminalError{Err: errors.New("something went wrong")},
			),
			expected: true,
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			actual := IsTerminal(testCase.err)
			require.Equal(t, testCase.expected, actual)
		})
	}
}
