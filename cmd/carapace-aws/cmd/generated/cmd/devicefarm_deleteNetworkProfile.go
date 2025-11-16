package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_deleteNetworkProfileCmd = &cobra.Command{
	Use:   "delete-network-profile",
	Short: "Deletes a network profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_deleteNetworkProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_deleteNetworkProfileCmd).Standalone()

		devicefarm_deleteNetworkProfileCmd.Flags().String("arn", "", "The ARN of the network profile to delete.")
		devicefarm_deleteNetworkProfileCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_deleteNetworkProfileCmd)
}
