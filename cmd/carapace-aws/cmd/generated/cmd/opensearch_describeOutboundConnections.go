package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_describeOutboundConnectionsCmd = &cobra.Command{
	Use:   "describe-outbound-connections",
	Short: "Lists all the outbound cross-cluster connections for a local (source) Amazon OpenSearch Service domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_describeOutboundConnectionsCmd).Standalone()

	opensearch_describeOutboundConnectionsCmd.Flags().String("filters", "", "List of filter names and values that you can use for requests.")
	opensearch_describeOutboundConnectionsCmd.Flags().String("max-results", "", "An optional parameter that specifies the maximum number of results to return.")
	opensearch_describeOutboundConnectionsCmd.Flags().String("next-token", "", "If your initial `DescribeOutboundConnections` operation returns a `nextToken`, you can include the returned `nextToken` in subsequent `DescribeOutboundConnections` operations, which returns results in the next page.")
	opensearchCmd.AddCommand(opensearch_describeOutboundConnectionsCmd)
}
