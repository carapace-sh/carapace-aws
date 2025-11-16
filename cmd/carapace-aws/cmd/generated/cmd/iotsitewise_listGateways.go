package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listGatewaysCmd = &cobra.Command{
	Use:   "list-gateways",
	Short: "Retrieves a paginated list of gateways.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listGatewaysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listGatewaysCmd).Standalone()

		iotsitewise_listGatewaysCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listGatewaysCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listGatewaysCmd)
}
