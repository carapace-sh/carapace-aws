package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var networkmanagerCmd = &cobra.Command{
	Use:   "networkmanager",
	Short: "Amazon Web Services enables you to centrally manage your Amazon Web Services Cloud WAN core network and your Transit Gateway network across Amazon Web Services accounts, Regions, and on-premises locations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(networkmanagerCmd).Standalone()

	rootCmd.AddCommand(networkmanagerCmd)
}
