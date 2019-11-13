package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestString(t *testing.T) {
	s := "a4bc2d5e"
	require.Equal(t, check_string(s), "aaaabccddddde", "must return string"+s)
}

func TestStringWithoutDigits(t *testing.T) {
	s := "abcd"
	require.Equal(t, check_string(s), "abcd", "must return self"+s)
}

func TestStringEmty(t *testing.T) {
	s := "45"
	require.Equal(t, check_string(s), "", "must return empty"+s)
}

func TestStringSlachOther(t *testing.T) {
	s := `qwe\45`
	require.Equal(t, check_string(s), `qwe44444`, "must return empty"+s)
}

func TestStringSlach(t *testing.T) {
	s := `qwe\\5`
	require.Equal(t, check_string(s), `qwe\\\\\`, "must return empty"+s)
}
