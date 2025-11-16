package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_describeOutboundCrossClusterSearchConnectionsCmd = &cobra.Command{
	Use:   "describe-outbound-cross-cluster-search-connections",
	Short: "Lists all the outbound cross-cluster search connections for a source domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_describeOutboundCrossClusterSearchConnectionsCmd).Standalone()

	es_describeOutboundCrossClusterSearchConnectionsCmd.Flags().String("filters", "", "A list of filters used to match properties for outbound cross-cluster search connection.")
	es_describeOutboundCrossClusterSearchConnectionsCmd.Flags().String("max-results", "", "Set this value to limit the number of results returned.")
	es_describeOutboundCrossClusterSearchConnectionsCmd.Flags().String("next-token", "", "NextToken is sent in case the earlier API call results contain the NextToken.")
	esCmd.AddCommand(es_describeOutboundCrossClusterSearchConnectionsCmd)
}
