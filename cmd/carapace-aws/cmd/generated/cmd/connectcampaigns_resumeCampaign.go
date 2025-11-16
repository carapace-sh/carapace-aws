package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_resumeCampaignCmd = &cobra.Command{
	Use:   "resume-campaign",
	Short: "Stops a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_resumeCampaignCmd).Standalone()

	connectcampaigns_resumeCampaignCmd.Flags().String("id", "", "")
	connectcampaigns_resumeCampaignCmd.MarkFlagRequired("id")
	connectcampaignsCmd.AddCommand(connectcampaigns_resumeCampaignCmd)
}
