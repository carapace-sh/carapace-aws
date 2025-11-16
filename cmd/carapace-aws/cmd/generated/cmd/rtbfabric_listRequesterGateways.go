package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rtbfabric_listRequesterGatewaysCmd = &cobra.Command{
	Use:   "list-requester-gateways",
	Short: "Lists requester gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rtbfabric_listRequesterGatewaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rtbfabric_listRequesterGatewaysCmd).Standalone()

		rtbfabric_listRequesterGatewaysCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		rtbfabric_listRequesterGatewaysCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	})
	rtbfabricCmd.AddCommand(rtbfabric_listRequesterGatewaysCmd)
}
