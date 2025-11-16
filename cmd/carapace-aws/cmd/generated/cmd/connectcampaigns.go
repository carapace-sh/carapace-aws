package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsCmd = &cobra.Command{
	Use:   "connectcampaigns",
	Short: "Provide APIs to create and manage Amazon Connect Campaigns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsCmd).Standalone()

	})
	rootCmd.AddCommand(connectcampaignsCmd)
}
