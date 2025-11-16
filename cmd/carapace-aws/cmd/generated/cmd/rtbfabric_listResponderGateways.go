package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_listResponderGatewaysCmd = &cobra.Command{
	Use:   "list-responder-gateways",
	Short: "Lists reponder gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_listResponderGatewaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_listResponderGatewaysCmd).Standalone()

		rtbfabric_listResponderGatewaysCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		rtbfabric_listResponderGatewaysCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	})
	rtbfabricCmd.AddCommand(rtbfabric_listResponderGatewaysCmd)
}
