package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_deleteDedicatedIpPoolCmd = &cobra.Command{
	Use:   "delete-dedicated-ip-pool",
	Short: "Delete a dedicated IP pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_deleteDedicatedIpPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_deleteDedicatedIpPoolCmd).Standalone()

		pinpointEmail_deleteDedicatedIpPoolCmd.Flags().String("pool-name", "", "The name of the dedicated IP pool that you want to delete.")
		pinpointEmail_deleteDedicatedIpPoolCmd.MarkFlagRequired("pool-name")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_deleteDedicatedIpPoolCmd)
}
