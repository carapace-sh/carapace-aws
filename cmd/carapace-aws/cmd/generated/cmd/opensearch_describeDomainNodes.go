package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeDomainNodesCmd = &cobra.Command{
	Use:   "describe-domain-nodes",
	Short: "Returns information about domain and nodes, including data nodes, master nodes, ultrawarm nodes, Availability Zone(s), standby nodes, node configurations, and node states.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeDomainNodesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_describeDomainNodesCmd).Standalone()

		opensearch_describeDomainNodesCmd.Flags().String("domain-name", "", "The name of the domain.")
		opensearch_describeDomainNodesCmd.MarkFlagRequired("domain-name")
	})
	opensearchCmd.AddCommand(opensearch_describeDomainNodesCmd)
}
