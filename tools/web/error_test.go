package web_test

import (
	"github.com/Sebalvarez97/mutants-go/tools/web"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewError(t *testing.T) {
	err := web.NewError(http.StatusBadRequest, "error occurred")
	require.Error(t, err)
	require.EqualValues(t, "400 bad_request: error occurred", err.Error())
}

func TestNewErrorf(t *testing.T) {
	err := web.NewErrorf(http.StatusBadRequest, "error occurred: %s", "detail")
	require.Error(t, err)
	require.EqualValues(t, "400 bad_request: error occurred: detail", err.Error())
}
