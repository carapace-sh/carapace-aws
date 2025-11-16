package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_describeCampaignCmd = &cobra.Command{
	Use:   "describe-campaign",
	Short: "Describes the specific campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_describeCampaignCmd).Standalone()

	connectcampaignsv2_describeCampaignCmd.Flags().String("id", "", "")
	connectcampaignsv2_describeCampaignCmd.MarkFlagRequired("id")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_describeCampaignCmd)
}
