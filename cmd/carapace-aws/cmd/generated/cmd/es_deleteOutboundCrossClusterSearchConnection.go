package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_deleteOutboundCrossClusterSearchConnectionCmd = &cobra.Command{
	Use:   "delete-outbound-cross-cluster-search-connection",
	Short: "Allows the source domain owner to delete an existing outbound cross-cluster search connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_deleteOutboundCrossClusterSearchConnectionCmd).Standalone()

	es_deleteOutboundCrossClusterSearchConnectionCmd.Flags().String("cross-cluster-search-connection-id", "", "The id of the outbound connection that you want to permanently delete.")
	es_deleteOutboundCrossClusterSearchConnectionCmd.MarkFlagRequired("cross-cluster-search-connection-id")
	esCmd.AddCommand(es_deleteOutboundCrossClusterSearchConnectionCmd)
}
