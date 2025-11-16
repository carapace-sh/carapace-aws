package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateNetworkCmd = &cobra.Command{
	Use:   "update-network",
	Short: "Change the settings for a Network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateNetworkCmd).Standalone()

	medialive_updateNetworkCmd.Flags().String("ip-pools", "", "Include this parameter only if you want to change the pool of IP addresses in the network.")
	medialive_updateNetworkCmd.Flags().String("name", "", "Include this parameter only if you want to change the name of the Network.")
	medialive_updateNetworkCmd.Flags().String("network-id", "", "The ID of the network")
	medialive_updateNetworkCmd.Flags().String("routes", "", "Include this parameter only if you want to change or add routes in the Network.")
	medialive_updateNetworkCmd.MarkFlagRequired("network-id")
	medialiveCmd.AddCommand(medialive_updateNetworkCmd)
}
