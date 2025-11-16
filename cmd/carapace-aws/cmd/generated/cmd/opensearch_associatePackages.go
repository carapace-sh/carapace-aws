package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_associatePackagesCmd = &cobra.Command{
	Use:   "associate-packages",
	Short: "Operation in the Amazon OpenSearch Service API for associating multiple packages with a domain simultaneously.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_associatePackagesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_associatePackagesCmd).Standalone()

		opensearch_associatePackagesCmd.Flags().String("domain-name", "", "")
		opensearch_associatePackagesCmd.Flags().String("package-list", "", "A list of packages and their prerequisites to be associated with a domain.")
		opensearch_associatePackagesCmd.MarkFlagRequired("domain-name")
		opensearch_associatePackagesCmd.MarkFlagRequired("package-list")
	})
	opensearchCmd.AddCommand(opensearch_associatePackagesCmd)
}
