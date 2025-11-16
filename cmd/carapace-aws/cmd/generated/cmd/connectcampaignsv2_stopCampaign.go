package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_stopCampaignCmd = &cobra.Command{
	Use:   "stop-campaign",
	Short: "Stops a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_stopCampaignCmd).Standalone()

	connectcampaignsv2_stopCampaignCmd.Flags().String("id", "", "")
	connectcampaignsv2_stopCampaignCmd.MarkFlagRequired("id")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_stopCampaignCmd)
}
