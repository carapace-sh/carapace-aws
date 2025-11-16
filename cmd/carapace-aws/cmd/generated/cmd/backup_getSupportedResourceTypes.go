package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_getSupportedResourceTypesCmd = &cobra.Command{
	Use:   "get-supported-resource-types",
	Short: "Returns the Amazon Web Services resource types supported by Backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_getSupportedResourceTypesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_getSupportedResourceTypesCmd).Standalone()

	})
	backupCmd.AddCommand(backup_getSupportedResourceTypesCmd)
}
