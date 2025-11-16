package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkFirewallCmd = &cobra.Command{
	Use:   "network-firewall",
	Short: "This is the API Reference for Network Firewall.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkFirewallCmd).Standalone()

	rootCmd.AddCommand(networkFirewallCmd)
}
