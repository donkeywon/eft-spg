package profile

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSvc_Open(t *testing.T) {
	svc := New(NewConfig())

	err := svc.Open()
	assert.NoError(t, err, "open save fail")
}
