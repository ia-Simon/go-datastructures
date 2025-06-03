package usecases_test

import (
	"testing"

	"github.com/ia-Simon/go-datastrutures/structures/stack/usecases"

	"github.com/stretchr/testify/require"
)

func TestInfixToPostfix(t *testing.T) {
	t.Run("flat operation", func(t *testing.T) {
		infix := "a+b*c+d"
		postfix, err := usecases.InfixToPostfix(infix)
		require.Nil(t, err)
		require.Equal(t, "abc*+d+", postfix)
	})

	t.Run("parenthesis operation", func(t *testing.T) {
		infix := "a+b*(c^d-e)^(f+g*h)-i"
		postfix, err := usecases.InfixToPostfix(infix)
		require.Nil(t, err)
		require.Equal(t, "abcd^e-fgh*+^*+i-", postfix)
	})

	t.Run("nested parenthesis operation", func(t *testing.T) {
		infix := "a^(b/(c+d))"
		postfix, err := usecases.InfixToPostfix(infix)
		require.Nil(t, err)
		require.Equal(t, "abcd+/^", postfix)
	})

	t.Run("unmatched open parenthesis", func(t *testing.T) {
		infix := "a^(b/(c+d)"
		_, err := usecases.InfixToPostfix(infix)
		require.Error(t, err)
	})

	t.Run("unexpected close parenthesis", func(t *testing.T) {
		infix := "a^b/(c+d))"
		_, err := usecases.InfixToPostfix(infix)
		require.Error(t, err)
	})
}
