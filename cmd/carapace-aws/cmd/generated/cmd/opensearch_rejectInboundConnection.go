package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_rejectInboundConnectionCmd = &cobra.Command{
	Use:   "reject-inbound-connection",
	Short: "Allows the remote Amazon OpenSearch Service domain owner to reject an inbound cross-cluster connection request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_rejectInboundConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_rejectInboundConnectionCmd).Standalone()

		opensearch_rejectInboundConnectionCmd.Flags().String("connection-id", "", "The unique identifier of the inbound connection to reject.")
		opensearch_rejectInboundConnectionCmd.MarkFlagRequired("connection-id")
	})
	opensearchCmd.AddCommand(opensearch_rejectInboundConnectionCmd)
}
