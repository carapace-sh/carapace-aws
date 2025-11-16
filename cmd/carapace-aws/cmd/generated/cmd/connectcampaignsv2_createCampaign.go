package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_createCampaignCmd = &cobra.Command{
	Use:   "create-campaign",
	Short: "Creates a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_createCampaignCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_createCampaignCmd).Standalone()

		connectcampaignsv2_createCampaignCmd.Flags().String("channel-subtype-config", "", "")
		connectcampaignsv2_createCampaignCmd.Flags().String("communication-limits-override", "", "")
		connectcampaignsv2_createCampaignCmd.Flags().String("communication-time-config", "", "")
		connectcampaignsv2_createCampaignCmd.Flags().String("connect-campaign-flow-arn", "", "")
		connectcampaignsv2_createCampaignCmd.Flags().String("connect-instance-id", "", "")
		connectcampaignsv2_createCampaignCmd.Flags().String("name", "", "")
		connectcampaignsv2_createCampaignCmd.Flags().String("schedule", "", "")
		connectcampaignsv2_createCampaignCmd.Flags().String("source", "", "")
		connectcampaignsv2_createCampaignCmd.Flags().String("tags", "", "")
		connectcampaignsv2_createCampaignCmd.MarkFlagRequired("channel-subtype-config")
		connectcampaignsv2_createCampaignCmd.MarkFlagRequired("connect-instance-id")
		connectcampaignsv2_createCampaignCmd.MarkFlagRequired("name")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_createCampaignCmd)
}
