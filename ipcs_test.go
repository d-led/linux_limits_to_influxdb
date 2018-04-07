package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// ipcs -ls

func TestParsingReferenceIpcsOutput(t *testing.T) {
	output := `
------ Semaphore Limits --------
max number of arrays = 32000
max semaphores per array = 32000
 max semaphores system wide  =1024000000
max ops per semop call = 500
semaphore max value = 32767
`
	res := parseIpcsLimits(output)

	assert.Equal(t, int64(32000), res["max_number_of_arrays"])
	assert.Equal(t, int64(1024000000), res["max_semaphores_system_wide"])
	assert.Equal(t, 5, len(res))
}

func TestSubstitutingSpacesInVariableNames(t *testing.T) {
	assert.Equal(t, "max_number_of_arrays", validIdentifierFrom("max number of arrays"), "spaces should be replaced with underscores")
	assert.Equal(t, "a_b", validIdentifierFrom(" a b  \t"), "surrounding spaces should be trimmed")
	assert.Equal(t, "a_b", validIdentifierFrom("a  \tb"), "multiple spaces should be replaced with only one underscore")
}
