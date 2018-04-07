package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// ipcs -ls
// /go/src/llti # ipcs -ls

func TestParsingIpcsOutpus(t *testing.T) {
	output := `
    ------ Semaphore Limits --------
max number of arrays = 32000
max semaphores per array = 32000
max semaphores system wide = 1024000000
max ops per semop call = 500
semaphore max value = 32767
`
	res := parseIpcsLimits(output)

	assert.Equal(t, 32000, res["max_number_of_arrays"])
}
