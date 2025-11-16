package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var opensearch_createOutboundConnectionCmd = &cobra.Command{
	Use:   "create-outbound-connection",
	Short: "Creates a new cross-cluster search connection from a source Amazon OpenSearch Service domain to a destination domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opensearch_createOutboundConnectionCmd).Standalone()

	opensearch_createOutboundConnectionCmd.Flags().String("connection-alias", "", "Name of the connection.")
	opensearch_createOutboundConnectionCmd.Flags().String("connection-mode", "", "The connection mode.")
	opensearch_createOutboundConnectionCmd.Flags().String("connection-properties", "", "The `ConnectionProperties` for the outbound connection.")
	opensearch_createOutboundConnectionCmd.Flags().String("local-domain-info", "", "Name and Region of the source (local) domain.")
	opensearch_createOutboundConnectionCmd.Flags().String("remote-domain-info", "", "Name and Region of the destination (remote) domain.")
	opensearch_createOutboundConnectionCmd.MarkFlagRequired("connection-alias")
	opensearch_createOutboundConnectionCmd.MarkFlagRequired("local-domain-info")
	opensearch_createOutboundConnectionCmd.MarkFlagRequired("remote-domain-info")
	opensearchCmd.AddCommand(opensearch_createOutboundConnectionCmd)
}
