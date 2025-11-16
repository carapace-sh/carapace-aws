package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_acceptInboundConnectionCmd = &cobra.Command{
	Use:   "accept-inbound-connection",
	Short: "Allows the destination Amazon OpenSearch Service domain owner to accept an inbound cross-cluster search connection request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_acceptInboundConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(opensearch_acceptInboundConnectionCmd).Standalone()

		opensearch_acceptInboundConnectionCmd.Flags().String("connection-id", "", "The ID of the inbound connection to accept.")
		opensearch_acceptInboundConnectionCmd.MarkFlagRequired("connection-id")
	})
	opensearchCmd.AddCommand(opensearch_acceptInboundConnectionCmd)
}
