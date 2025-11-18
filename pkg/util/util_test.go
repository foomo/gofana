package util_test

import (
	"testing"

	"github.com/foomo/gofana/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestMustYamlToMap(t *testing.T) {
	t.Parallel()

	s := `
		elimiter: ","
		format: json
		jsonPaths:
		- alias: ""
		  path: squadron
		- alias: ""
		  path: user
		- alias: ""
		  path: commit
		- alias: ""
		  path: branch
		keepTime: false
		replace: false
		source: description
	`
	o := util.MustYamlToMap(s)
	assert.Len(t, o, 6)
}
