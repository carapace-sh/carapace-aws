package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createNetworkCmd = &cobra.Command{
	Use:   "create-network",
	Short: "Create as many Networks as you need.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_createNetworkCmd).Standalone()

		medialive_createNetworkCmd.Flags().String("ip-pools", "", "An array of IpPoolCreateRequests that identify a collection of IP addresses in your network that you want to reserve for use in MediaLive Anywhere.")
		medialive_createNetworkCmd.Flags().String("name", "", "Specify a name that is unique in the AWS account.")
		medialive_createNetworkCmd.Flags().String("request-id", "", "An ID that you assign to a create request.")
		medialive_createNetworkCmd.Flags().String("routes", "", "An array of routes that MediaLive Anywhere needs to know about in order to route encoding traffic.")
		medialive_createNetworkCmd.Flags().String("tags", "", "A collection of key-value pairs.")
	})
	medialiveCmd.AddCommand(medialive_createNetworkCmd)
}
