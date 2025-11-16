package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var es_createOutboundCrossClusterSearchConnectionCmd = &cobra.Command{
	Use:   "create-outbound-cross-cluster-search-connection",
	Short: "Creates a new cross-cluster search connection from a source domain to a destination domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(es_createOutboundCrossClusterSearchConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(es_createOutboundCrossClusterSearchConnectionCmd).Standalone()

		es_createOutboundCrossClusterSearchConnectionCmd.Flags().String("connection-alias", "", "Specifies the connection alias that will be used by the customer for this connection.")
		es_createOutboundCrossClusterSearchConnectionCmd.Flags().String("destination-domain-info", "", "Specifies the `DomainInformation` for the destination Elasticsearch domain.")
		es_createOutboundCrossClusterSearchConnectionCmd.Flags().String("source-domain-info", "", "Specifies the `DomainInformation` for the source Elasticsearch domain.")
		es_createOutboundCrossClusterSearchConnectionCmd.MarkFlagRequired("connection-alias")
		es_createOutboundCrossClusterSearchConnectionCmd.MarkFlagRequired("destination-domain-info")
		es_createOutboundCrossClusterSearchConnectionCmd.MarkFlagRequired("source-domain-info")
	})
	esCmd.AddCommand(es_createOutboundCrossClusterSearchConnectionCmd)
}
