package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// cat /etc/*-release

func TestParsingReleaseContents(t *testing.T) {
	alpine := `
3.7.0
NAME="Alpine Linux"
ID=alpine
VERSION_ID=3.7.0
PRETTY_NAME="Alpine Linux v3.7"
HOME_URL="http://alpinelinux.org"
BUG_REPORT_URL="http://bugs.alpinelinux.org"
`
	res := parseRelease(alpine)

	assert.Equal(t, "3.7.0", res["VERSION_ID"])
	assert.Equal(t, "alpine", res["ID"])
	assert.Equal(t, "Alpine Linux v3.7", res["PRETTY_NAME"], "quotes should be trimmed away")
	assert.Equal(t, 6, len(res))

	debian := `
PRETTY_NAME="Debian GNU/Linux 9 (stretch)"
NAME="Debian GNU/Linux"
VERSION_ID="9"
VERSION="9 (stretch)"
ID=debian
HOME_URL="https://www.debian.org/"
SUPPORT_URL="https://www.debian.org/support"
BUG_REPORT_URL="https://bugs.debian.org/"
`

	res = parseRelease(debian)
	assert.Equal(t, "9", res["VERSION_ID"])
	assert.Equal(t, "debian", res["ID"])

}
