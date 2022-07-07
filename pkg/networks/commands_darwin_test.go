package networks

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestVDESock(t *testing.T) {
	config, err := DefaultConfig()
	assert.NilError(t, err)

	vdeSock := config.VDESock("foo")
	assert.Equal(t, vdeSock, "/private/var/run/lima/foo.ctl")
}

func TestPIDFile(t *testing.T) {
	config, err := DefaultConfig()
	assert.NilError(t, err)

	pidFile := config.PIDFile("name", "daemon")
	assert.Equal(t, pidFile, "/private/var/run/lima/name_daemon.pid")
}
