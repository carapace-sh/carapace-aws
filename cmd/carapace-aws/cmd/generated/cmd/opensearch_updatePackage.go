package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_updatePackageCmd = &cobra.Command{
	Use:   "update-package",
	Short: "Updates a package for use with Amazon OpenSearch Service domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_updatePackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_updatePackageCmd).Standalone()

		opensearch_updatePackageCmd.Flags().String("commit-message", "", "Commit message for the updated file, which is shown as part of `GetPackageVersionHistoryResponse`.")
		opensearch_updatePackageCmd.Flags().String("package-configuration", "", "The updated configuration details for a package.")
		opensearch_updatePackageCmd.Flags().String("package-description", "", "A new description of the package.")
		opensearch_updatePackageCmd.Flags().String("package-encryption-options", "", "Encryption options for a package.")
		opensearch_updatePackageCmd.Flags().String("package-id", "", "The unique identifier for the package.")
		opensearch_updatePackageCmd.Flags().String("package-source", "", "Amazon S3 bucket and key for the package.")
		opensearch_updatePackageCmd.MarkFlagRequired("package-id")
		opensearch_updatePackageCmd.MarkFlagRequired("package-source")
	})
	opensearchCmd.AddCommand(opensearch_updatePackageCmd)
}
