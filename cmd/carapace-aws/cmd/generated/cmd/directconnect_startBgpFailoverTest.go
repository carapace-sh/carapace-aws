package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_startBgpFailoverTestCmd = &cobra.Command{
	Use:   "start-bgp-failover-test",
	Short: "Starts the virtual interface failover test that verifies your configuration meets your resiliency requirements by placing the BGP peering session in the DOWN state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_startBgpFailoverTestCmd).Standalone()

	directconnect_startBgpFailoverTestCmd.Flags().String("bgp-peers", "", "The BGP peers to place in the DOWN state.")
	directconnect_startBgpFailoverTestCmd.Flags().String("test-duration-in-minutes", "", "The time in minutes that the virtual interface failover test will last.")
	directconnect_startBgpFailoverTestCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface you want to test.")
	directconnect_startBgpFailoverTestCmd.MarkFlagRequired("virtual-interface-id")
	directconnectCmd.AddCommand(directconnect_startBgpFailoverTestCmd)
}
