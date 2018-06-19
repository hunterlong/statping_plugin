package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMainUnits(t *testing.T) {
	info := Plugin.GetInfo()
	assert.Equal(t, "Example Plugin", info.Name)
	assert.Equal(t, "This is an example plugin for Statup Status Page application. It can be implemented pretty quick!", info.Description)
}
