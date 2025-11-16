package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_deleteOutboundConnectionCmd = &cobra.Command{
	Use:   "delete-outbound-connection",
	Short: "Allows the source Amazon OpenSearch Service domain owner to delete an existing outbound cross-cluster search connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_deleteOutboundConnectionCmd).Standalone()

	opensearch_deleteOutboundConnectionCmd.Flags().String("connection-id", "", "The ID of the outbound connection you want to permanently delete.")
	opensearch_deleteOutboundConnectionCmd.MarkFlagRequired("connection-id")
	opensearchCmd.AddCommand(opensearch_deleteOutboundConnectionCmd)
}
