package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_describeFrameworkCmd = &cobra.Command{
	Use:   "describe-framework",
	Short: "Returns the framework details for the specified `FrameworkName`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_describeFrameworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_describeFrameworkCmd).Standalone()

		backup_describeFrameworkCmd.Flags().String("framework-name", "", "The unique name of a framework.")
		backup_describeFrameworkCmd.MarkFlagRequired("framework-name")
	})
	backupCmd.AddCommand(backup_describeFrameworkCmd)
}
