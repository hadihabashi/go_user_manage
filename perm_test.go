package go_user_manage
import (
	"testing"
)

func TestInterface(t *testing.T) {
	// Check that the value qualifies for the interface
	perm,err := New(0,"","")
	if err!= nil {
		t.Error("Error, " + err.Error())
	}
	var _ IPermissions = perm
}

