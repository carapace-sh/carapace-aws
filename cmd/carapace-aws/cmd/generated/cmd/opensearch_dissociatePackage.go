package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_dissociatePackageCmd = &cobra.Command{
	Use:   "dissociate-package",
	Short: "Removes a package from the specified Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_dissociatePackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_dissociatePackageCmd).Standalone()

		opensearch_dissociatePackageCmd.Flags().String("domain-name", "", "Name of the domain to dissociate the package from.")
		opensearch_dissociatePackageCmd.Flags().String("package-id", "", "Internal ID of the package to dissociate from the domain.")
		opensearch_dissociatePackageCmd.MarkFlagRequired("domain-name")
		opensearch_dissociatePackageCmd.MarkFlagRequired("package-id")
	})
	opensearchCmd.AddCommand(opensearch_dissociatePackageCmd)
}
