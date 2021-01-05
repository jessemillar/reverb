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

func TestTextArgumentMultipleLines(t *testing.T) {
	var b bytes.Buffer

	var conf argConfig
	conf.separator = "-"
	conf.enableEscapeSequences = true
	conf.args = []string{"Test\nTest"}

	reverb(width, &conf, &b)
	assert.Equal(t, "---\nTest\nTest\n---\n", b.String(), "Unexpected output")
}

func TestTextArgumentMultipleLinesLong(t *testing.T) {
	var b bytes.Buffer

	var conf argConfig
	conf.separator = "-"
	conf.enableEscapeSequences = true
	conf.args = []string{"Test\nTest test test test test test test test test test test test"}

	// Pass 0 to simulate a headless environment
	reverb(0, &conf, &b)
	assert.Equal(t, "-----------------------------------------------------------\nTest\nTest test test test test test test test test test test test\n-----------------------------------------------------------\n", b.String(), "Unexpected output")
}

func TestTextArgumentWithEscapeSequences(t *testing.T) {
	var b bytes.Buffer

	var conf argConfig
	conf.separator = "-"
	conf.args = []string{"Test\n\ttest\ntest"}

	reverb(width, &conf, &b)
	assert.Equal(t, "---\nTest\n\ttest\ntest\n---\n", b.String(), "Unexpected output")
}

func TestDynamicSeparatorWidth(t *testing.T) {
	var b bytes.Buffer

	var conf argConfig
	conf.separator = "-"
	conf.args = []string{"Test test test test test test test test test test"}

	// Pass 0 to simulate a headless environment
	reverb(0, &conf, &b)
	assert.Equal(t, "-------------------------------------------------\nTest test test test test test test test test test\n-------------------------------------------------\n", b.String(), "Unexpected output")
}

func TestStaticSeparatorWidth(t *testing.T) {
	var b bytes.Buffer

	var conf argConfig
	conf.separator = "-"
	conf.disableDynamicWidth = true
	conf.args = []string{"Test test test test test test test test test test"}

	// Pass 0 to simulate a headless environment
	reverb(0, &conf, &b)
	assert.Equal(t, "------------------------------------------\nTest test test test test test test test test test\n------------------------------------------\n", b.String(), "Unexpected output")
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

func TestLongCustomSeparator(t *testing.T) {
	var b bytes.Buffer

	var conf argConfig
	// Should only grab the first character from the string
	conf.separator = "~-+"

	reverb(width, &conf, &b)
	assert.Equal(t, "Please pass only one character to the -c flag\n", b.String(), "Unexpected output")
}
