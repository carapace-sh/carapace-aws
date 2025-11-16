package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_listGatewaysCmd = &cobra.Command{
	Use:   "list-gateways",
	Short: "Displays a list of gateways that are associated with this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_listGatewaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_listGatewaysCmd).Standalone()

		mediaconnect_listGatewaysCmd.Flags().String("max-results", "", "The maximum number of results to return per API request.")
		mediaconnect_listGatewaysCmd.Flags().String("next-token", "", "The token that identifies the batch of results that you want to see.")
	})
	mediaconnectCmd.AddCommand(mediaconnect_listGatewaysCmd)
}
