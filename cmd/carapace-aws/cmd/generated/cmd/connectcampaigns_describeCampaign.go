package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_describeCampaignCmd = &cobra.Command{
	Use:   "describe-campaign",
	Short: "Describes the specific campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_describeCampaignCmd).Standalone()

	connectcampaigns_describeCampaignCmd.Flags().String("id", "", "")
	connectcampaigns_describeCampaignCmd.MarkFlagRequired("id")
	connectcampaignsCmd.AddCommand(connectcampaigns_describeCampaignCmd)
}
