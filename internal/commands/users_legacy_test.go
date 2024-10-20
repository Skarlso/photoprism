package commands

import (
	"testing"

	"github.com/photoprism/photoprism/pkg/capture"
	"github.com/stretchr/testify/assert"
)

func TestUsersLegacyCommand(t *testing.T) {
	t.Run("All", func(t *testing.T) {
		var err error

		// Create test context with flags and arguments.
		ctx := NewTestContext([]string{""})

		// Run command with test context.
		output := capture.Output(func() {
			err = UsersLegacyCommand.Run(ctx)
		})

		// Check command output for plausibility.
		//t.Logf(output)
		assert.NoError(t, err)
		assert.Contains(t, output, "| ID | UID | Name | User | Email | Admin | Created At |")
	})
}
