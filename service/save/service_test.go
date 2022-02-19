package save

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSvc_Open(t *testing.T) {
	svc := New(NewConfig())

	err := svc.Open()
	assert.NoError(t, err, "open save fail")
	fmt.Println(os.ModeExclusive.String())
}
