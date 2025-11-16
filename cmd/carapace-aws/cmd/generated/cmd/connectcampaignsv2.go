package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2Cmd = &cobra.Command{
	Use:   "connectcampaignsv2",
	Short: "Provide APIs to create and manage Amazon Connect Campaigns.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2Cmd).Standalone()

	})
	rootCmd.AddCommand(connectcampaignsv2Cmd)
}
