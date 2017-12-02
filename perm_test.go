package go_user_manage
import (
	"testing"
)

func TestInterface(t *testing.T) {
	// Check that the value qualifies for the interface
	var _ IPermissions = New()
}

