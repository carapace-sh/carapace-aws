package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanager_createGlobalNetworkCmd = &cobra.Command{
	Use:   "create-global-network",
	Short: "Creates a new, empty global network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanager_createGlobalNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(networkmanager_createGlobalNetworkCmd).Standalone()

		networkmanager_createGlobalNetworkCmd.Flags().String("description", "", "A description of the global network.")
		networkmanager_createGlobalNetworkCmd.Flags().String("tags", "", "The tags to apply to the resource during creation.")
	})
	networkmanagerCmd.AddCommand(networkmanager_createGlobalNetworkCmd)
}
