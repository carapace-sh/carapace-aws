package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_rejectInboundCrossClusterSearchConnectionCmd = &cobra.Command{
	Use:   "reject-inbound-cross-cluster-search-connection",
	Short: "Allows the destination domain owner to reject an inbound cross-cluster search connection request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_rejectInboundCrossClusterSearchConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_rejectInboundCrossClusterSearchConnectionCmd).Standalone()

		es_rejectInboundCrossClusterSearchConnectionCmd.Flags().String("cross-cluster-search-connection-id", "", "The id of the inbound connection that you want to reject.")
		es_rejectInboundCrossClusterSearchConnectionCmd.MarkFlagRequired("cross-cluster-search-connection-id")
	})
	esCmd.AddCommand(es_rejectInboundCrossClusterSearchConnectionCmd)
}
