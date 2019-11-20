package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestString(t *testing.T) {
	s := "dfda dff dasf dfsf dfad ds ds ds ffe fef"
	require.Equal(t, unique(s), [{d 4}, {man 3}, {dfs 1}, {dfsd 1} ], "must return string"+s)
}

