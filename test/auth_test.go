package test

import (
	"testing"

	"github.com/RoyJoel/LZStudyingProjectBackendBackend/package/web/auth"
)

func TestAuth(t *testing.T) {
	auth.DeletePolicy("1", "", "*")
}
