package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMixCase(t *testing.T) {
	require.Equal(t, "PrAcTiCe — ТеСтОвОе", MixCase("Practice — тестовое"))
	require.Empty(t, MixCase(""))
}
