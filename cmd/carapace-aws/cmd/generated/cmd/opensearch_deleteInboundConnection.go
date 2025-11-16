package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_deleteInboundConnectionCmd = &cobra.Command{
	Use:   "delete-inbound-connection",
	Short: "Allows the destination Amazon OpenSearch Service domain owner to delete an existing inbound cross-cluster search connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_deleteInboundConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_deleteInboundConnectionCmd).Standalone()

		opensearch_deleteInboundConnectionCmd.Flags().String("connection-id", "", "The ID of the inbound connection to permanently delete.")
		opensearch_deleteInboundConnectionCmd.MarkFlagRequired("connection-id")
	})
	opensearchCmd.AddCommand(opensearch_deleteInboundConnectionCmd)
}
