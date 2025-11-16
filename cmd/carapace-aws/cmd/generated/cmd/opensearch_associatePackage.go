package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_associatePackageCmd = &cobra.Command{
	Use:   "associate-package",
	Short: "Associates a package with an Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_associatePackageCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_associatePackageCmd).Standalone()

		opensearch_associatePackageCmd.Flags().String("association-configuration", "", "The configuration for associating a package with an Amazon OpenSearch Service domain.")
		opensearch_associatePackageCmd.Flags().String("domain-name", "", "Name of the domain to associate the package with.")
		opensearch_associatePackageCmd.Flags().String("package-id", "", "Internal ID of the package to associate with a domain.")
		opensearch_associatePackageCmd.Flags().String("prerequisite-package-idlist", "", "A list of package IDs that must be associated with the domain before the package specified in the request can be associated.")
		opensearch_associatePackageCmd.MarkFlagRequired("domain-name")
		opensearch_associatePackageCmd.MarkFlagRequired("package-id")
	})
	opensearchCmd.AddCommand(opensearch_associatePackageCmd)
}
