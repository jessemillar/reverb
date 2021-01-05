package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const width = 3

func TestVersion(t *testing.T) {
	var b bytes.Buffer

	var conf argConfig
	conf.separator = "-"
	conf.version = true

	reverb(width, &conf, &b)
	assert.Equal(t, version+"\n", b.String(), "Version numbers didn't match")
}

func TestNoTextArgument(t *testing.T) {
	var b bytes.Buffer

	var conf argConfig
	conf.separator = "-"

	reverb(width, &conf, &b)
	assert.Equal(t, "---\n", b.String(), "Unexpected output")
}

func TestTextArgument(t *testing.T) {
	var b bytes.Buffer

	var conf argConfig
	conf.separator = "-"
	conf.args = []string{"Test"}

	reverb(width, &conf, &b)
	assert.Equal(t, "---\nTest\n---\n", b.String(), "Unexpected output")
}

func TestMultipleTextArguments(t *testing.T) {
	var b bytes.Buffer

	var conf argConfig
	conf.separator = "-"
	conf.args = []string{"Test", "test", "testing"}

	reverb(width, &conf, &b)
	assert.Equal(t, "---\nTest test testing\n---\n", b.String(), "Unexpected output")
}

func TestCustomSeparator(t *testing.T) {
	var b bytes.Buffer

	var conf argConfig
	conf.separator = "!"

	reverb(width, &conf, &b)
	assert.Equal(t, "!!!\n", b.String(), "Unexpected output")
}
