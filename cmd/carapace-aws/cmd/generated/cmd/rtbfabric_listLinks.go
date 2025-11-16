package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_listLinksCmd = &cobra.Command{
	Use:   "list-links",
	Short: "Lists links associated with gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_listLinksCmd).Standalone()

	rtbfabric_listLinksCmd.Flags().String("gateway-id", "", "The unique identifier of the gateway.")
	rtbfabric_listLinksCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
	rtbfabric_listLinksCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	rtbfabric_listLinksCmd.MarkFlagRequired("gateway-id")
	rtbfabricCmd.AddCommand(rtbfabric_listLinksCmd)
}
