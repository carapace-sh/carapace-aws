package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeDomainsCmd = &cobra.Command{
	Use:   "describe-domains",
	Short: "Returns domain configuration information about the specified Amazon OpenSearch Service domains.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeDomainsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_describeDomainsCmd).Standalone()

		opensearch_describeDomainsCmd.Flags().String("domain-names", "", "Array of OpenSearch Service domain names that you want information about.")
		opensearch_describeDomainsCmd.MarkFlagRequired("domain-names")
	})
	opensearchCmd.AddCommand(opensearch_describeDomainsCmd)
}
