package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listMultiRegionEndpointsCmd = &cobra.Command{
	Use:   "list-multi-region-endpoints",
	Short: "List the multi-region endpoints (global-endpoints).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listMultiRegionEndpointsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_listMultiRegionEndpointsCmd).Standalone()

		sesv2_listMultiRegionEndpointsCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListMultiRegionEndpoints` to indicate the position in the list of multi-region endpoints (global-endpoints).")
		sesv2_listMultiRegionEndpointsCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListMultiRegionEndpoints`.")
	})
	sesv2Cmd.AddCommand(sesv2_listMultiRegionEndpointsCmd)
}
