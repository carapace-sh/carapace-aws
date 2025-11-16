package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeInboundConnectionsCmd = &cobra.Command{
	Use:   "describe-inbound-connections",
	Short: "Lists all the inbound cross-cluster search connections for a destination (remote) Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeInboundConnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_describeInboundConnectionsCmd).Standalone()

		opensearch_describeInboundConnectionsCmd.Flags().String("filters", "", "A list of filters used to match properties for inbound cross-cluster connections.")
		opensearch_describeInboundConnectionsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
		opensearch_describeInboundConnectionsCmd.Flags().String("next-token", "", "If your initial `DescribeInboundConnections` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `DescribeInboundConnections` operations, which returns results in the next page.")
	})
	opensearchCmd.AddCommand(opensearch_describeInboundConnectionsCmd)
}
