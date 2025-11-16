package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_createFrameworkCmd = &cobra.Command{
	Use:   "create-framework",
	Short: "Creates a framework with one or more controls.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_createFrameworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_createFrameworkCmd).Standalone()

		backup_createFrameworkCmd.Flags().String("framework-controls", "", "The controls that make up the framework.")
		backup_createFrameworkCmd.Flags().String("framework-description", "", "An optional description of the framework with a maximum of 1,024 characters.")
		backup_createFrameworkCmd.Flags().String("framework-name", "", "The unique name of the framework.")
		backup_createFrameworkCmd.Flags().String("framework-tags", "", "The tags to assign to the framework.")
		backup_createFrameworkCmd.Flags().String("idempotency-token", "", "A customer-chosen string that you can use to distinguish between otherwise identical calls to `CreateFrameworkInput`.")
		backup_createFrameworkCmd.MarkFlagRequired("framework-controls")
		backup_createFrameworkCmd.MarkFlagRequired("framework-name")
	})
	backupCmd.AddCommand(backup_createFrameworkCmd)
}
