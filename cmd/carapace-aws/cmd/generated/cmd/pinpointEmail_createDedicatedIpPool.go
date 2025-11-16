package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_createDedicatedIpPoolCmd = &cobra.Command{
	Use:   "create-dedicated-ip-pool",
	Short: "Create a new pool of dedicated IP addresses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_createDedicatedIpPoolCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pinpointEmail_createDedicatedIpPoolCmd).Standalone()

		pinpointEmail_createDedicatedIpPoolCmd.Flags().String("pool-name", "", "The name of the dedicated IP pool.")
		pinpointEmail_createDedicatedIpPoolCmd.Flags().String("tags", "", "An object that defines the tags (keys and values) that you want to associate with the pool.")
		pinpointEmail_createDedicatedIpPoolCmd.MarkFlagRequired("pool-name")
	})
	pinpointEmailCmd.AddCommand(pinpointEmail_createDedicatedIpPoolCmd)
}
