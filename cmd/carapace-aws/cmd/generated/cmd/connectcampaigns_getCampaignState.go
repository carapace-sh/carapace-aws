package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_getCampaignStateCmd = &cobra.Command{
	Use:   "get-campaign-state",
	Short: "Get state of a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_getCampaignStateCmd).Standalone()

	connectcampaigns_getCampaignStateCmd.Flags().String("id", "", "")
	connectcampaigns_getCampaignStateCmd.MarkFlagRequired("id")
	connectcampaignsCmd.AddCommand(connectcampaigns_getCampaignStateCmd)
}
