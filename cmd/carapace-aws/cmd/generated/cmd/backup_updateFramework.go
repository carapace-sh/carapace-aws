package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_updateFrameworkCmd = &cobra.Command{
	Use:   "update-framework",
	Short: "Updates the specified framework.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_updateFrameworkCmd).Standalone()

	backup_updateFrameworkCmd.Flags().String("framework-controls", "", "The controls that make up the framework.")
	backup_updateFrameworkCmd.Flags().String("framework-description", "", "An optional description of the framework with a maximum 1,024 characters.")
	backup_updateFrameworkCmd.Flags().String("framework-name", "", "The unique name of a framework.")
	backup_updateFrameworkCmd.Flags().String("idempotency-token", "", "A customer-chosen string that you can use to distinguish between otherwise identical calls to `UpdateFrameworkInput`.")
	backup_updateFrameworkCmd.MarkFlagRequired("framework-name")
	backupCmd.AddCommand(backup_updateFrameworkCmd)
}
