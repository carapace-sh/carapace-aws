package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_pauseCampaignCmd = &cobra.Command{
	Use:   "pause-campaign",
	Short: "Pauses a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_pauseCampaignCmd).Standalone()

	connectcampaigns_pauseCampaignCmd.Flags().String("id", "", "")
	connectcampaigns_pauseCampaignCmd.MarkFlagRequired("id")
	connectcampaignsCmd.AddCommand(connectcampaigns_pauseCampaignCmd)
}
