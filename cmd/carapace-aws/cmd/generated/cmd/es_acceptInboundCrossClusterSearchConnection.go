package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_acceptInboundCrossClusterSearchConnectionCmd = &cobra.Command{
	Use:   "accept-inbound-cross-cluster-search-connection",
	Short: "Allows the destination domain owner to accept an inbound cross-cluster search connection request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_acceptInboundCrossClusterSearchConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_acceptInboundCrossClusterSearchConnectionCmd).Standalone()

		es_acceptInboundCrossClusterSearchConnectionCmd.Flags().String("cross-cluster-search-connection-id", "", "The id of the inbound connection that you want to accept.")
		es_acceptInboundCrossClusterSearchConnectionCmd.MarkFlagRequired("cross-cluster-search-connection-id")
	})
	esCmd.AddCommand(es_acceptInboundCrossClusterSearchConnectionCmd)
}
