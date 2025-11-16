package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describeInboundCrossClusterSearchConnectionsCmd = &cobra.Command{
	Use:   "describe-inbound-cross-cluster-search-connections",
	Short: "Lists all the inbound cross-cluster search connections for a destination domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describeInboundCrossClusterSearchConnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_describeInboundCrossClusterSearchConnectionsCmd).Standalone()

		es_describeInboundCrossClusterSearchConnectionsCmd.Flags().String("filters", "", "A list of filters used to match properties for inbound cross-cluster search connection.")
		es_describeInboundCrossClusterSearchConnectionsCmd.Flags().String("max-results", "", "Set this value to limit the number of results returned.")
		es_describeInboundCrossClusterSearchConnectionsCmd.Flags().String("next-token", "", "NextToken is sent in case the earlier API call results contain the NextToken.")
	})
	esCmd.AddCommand(es_describeInboundCrossClusterSearchConnectionsCmd)
}
