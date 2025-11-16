package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_createPackageCmd = &cobra.Command{
	Use:   "create-package",
	Short: "Creates a package for use with Amazon OpenSearch Service domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_createPackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_createPackageCmd).Standalone()

		opensearch_createPackageCmd.Flags().String("engine-version", "", "The version of the Amazon OpenSearch Service engine for which is compatible with the package.")
		opensearch_createPackageCmd.Flags().String("package-configuration", "", "The configuration parameters for the package being created.")
		opensearch_createPackageCmd.Flags().String("package-description", "", "Description of the package.")
		opensearch_createPackageCmd.Flags().String("package-encryption-options", "", "The encryption parameters for the package being created.")
		opensearch_createPackageCmd.Flags().String("package-name", "", "Unique name for the package.")
		opensearch_createPackageCmd.Flags().String("package-source", "", "The Amazon S3 location from which to import the package.")
		opensearch_createPackageCmd.Flags().String("package-type", "", "The type of package.")
		opensearch_createPackageCmd.Flags().String("package-vending-options", "", "The vending options for the package being created.")
		opensearch_createPackageCmd.MarkFlagRequired("package-name")
		opensearch_createPackageCmd.MarkFlagRequired("package-source")
		opensearch_createPackageCmd.MarkFlagRequired("package-type")
	})
	opensearchCmd.AddCommand(opensearch_createPackageCmd)
}
