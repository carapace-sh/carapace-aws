package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_updateCampaignScheduleCmd = &cobra.Command{
	Use:   "update-campaign-schedule",
	Short: "Updates the schedule for a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_updateCampaignScheduleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_updateCampaignScheduleCmd).Standalone()

		connectcampaignsv2_updateCampaignScheduleCmd.Flags().String("id", "", "")
		connectcampaignsv2_updateCampaignScheduleCmd.Flags().String("schedule", "", "")
		connectcampaignsv2_updateCampaignScheduleCmd.MarkFlagRequired("id")
		connectcampaignsv2_updateCampaignScheduleCmd.MarkFlagRequired("schedule")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_updateCampaignScheduleCmd)
}
