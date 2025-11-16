package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeDomainCmd = &cobra.Command{
	Use:   "describe-domain",
	Short: "Describes the domain configuration for the specified Amazon OpenSearch Service domain, including the domain ID, domain service endpoint, and domain ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeDomainCmd).Standalone()

	opensearch_describeDomainCmd.Flags().String("domain-name", "", "The name of the domain that you want information about.")
	opensearch_describeDomainCmd.MarkFlagRequired("domain-name")
	opensearchCmd.AddCommand(opensearch_describeDomainCmd)
}
