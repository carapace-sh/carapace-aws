package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_listDedicatedIpPoolsCmd = &cobra.Command{
	Use:   "list-dedicated-ip-pools",
	Short: "List all of the dedicated IP pools that exist in your Amazon Pinpoint account in the current AWS Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_listDedicatedIpPoolsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_listDedicatedIpPoolsCmd).Standalone()

		pinpointEmail_listDedicatedIpPoolsCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListDedicatedIpPools` to indicate the position in the list of dedicated IP pools.")
		pinpointEmail_listDedicatedIpPoolsCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListDedicatedIpPools`.")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_listDedicatedIpPoolsCmd)
}
