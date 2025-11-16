package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_deleteInboundCrossClusterSearchConnectionCmd = &cobra.Command{
	Use:   "delete-inbound-cross-cluster-search-connection",
	Short: "Allows the destination domain owner to delete an existing inbound cross-cluster search connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_deleteInboundCrossClusterSearchConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_deleteInboundCrossClusterSearchConnectionCmd).Standalone()

		es_deleteInboundCrossClusterSearchConnectionCmd.Flags().String("cross-cluster-search-connection-id", "", "The id of the inbound connection that you want to permanently delete.")
		es_deleteInboundCrossClusterSearchConnectionCmd.MarkFlagRequired("cross-cluster-search-connection-id")
	})
	esCmd.AddCommand(es_deleteInboundCrossClusterSearchConnectionCmd)
}
