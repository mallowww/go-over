//go:build !integration || !linux

package tag

import "testing"

func TestTagNOT(t *testing.T) {
	t.Log("tag integration not db")
}
