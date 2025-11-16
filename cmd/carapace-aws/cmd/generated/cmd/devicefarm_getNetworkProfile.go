package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getNetworkProfileCmd = &cobra.Command{
	Use:   "get-network-profile",
	Short: "Returns information about a network profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getNetworkProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_getNetworkProfileCmd).Standalone()

		devicefarm_getNetworkProfileCmd.Flags().String("arn", "", "The ARN of the network profile to return information about.")
		devicefarm_getNetworkProfileCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_getNetworkProfileCmd)
}
