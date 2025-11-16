package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_updatePackageScopeCmd = &cobra.Command{
	Use:   "update-package-scope",
	Short: "Updates the scope of a package.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_updatePackageScopeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_updatePackageScopeCmd).Standalone()

		opensearch_updatePackageScopeCmd.Flags().String("operation", "", "The operation to perform on the package scope (e.g., add/remove/override users).")
		opensearch_updatePackageScopeCmd.Flags().String("package-id", "", "ID of the package whose scope is being updated.")
		opensearch_updatePackageScopeCmd.Flags().String("package-user-list", "", "List of users to be added or removed from the package scope.")
		opensearch_updatePackageScopeCmd.MarkFlagRequired("operation")
		opensearch_updatePackageScopeCmd.MarkFlagRequired("package-id")
		opensearch_updatePackageScopeCmd.MarkFlagRequired("package-user-list")
	})
	opensearchCmd.AddCommand(opensearch_updatePackageScopeCmd)
}
