package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_stopBgpFailoverTestCmd = &cobra.Command{
	Use:   "stop-bgp-failover-test",
	Short: "Stops the virtual interface failover test.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_stopBgpFailoverTestCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_stopBgpFailoverTestCmd).Standalone()

		directconnect_stopBgpFailoverTestCmd.Flags().String("virtual-interface-id", "", "The ID of the virtual interface you no longer want to test.")
		directconnect_stopBgpFailoverTestCmd.MarkFlagRequired("virtual-interface-id")
	})
	directconnectCmd.AddCommand(directconnect_stopBgpFailoverTestCmd)
}
