package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_updateCampaignOutboundCallConfigCmd = &cobra.Command{
	Use:   "update-campaign-outbound-call-config",
	Short: "Updates the outbound call config of a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_updateCampaignOutboundCallConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaigns_updateCampaignOutboundCallConfigCmd).Standalone()

		connectcampaigns_updateCampaignOutboundCallConfigCmd.Flags().String("answer-machine-detection-config", "", "")
		connectcampaigns_updateCampaignOutboundCallConfigCmd.Flags().String("connect-contact-flow-id", "", "")
		connectcampaigns_updateCampaignOutboundCallConfigCmd.Flags().String("connect-source-phone-number", "", "")
		connectcampaigns_updateCampaignOutboundCallConfigCmd.Flags().String("id", "", "")
		connectcampaigns_updateCampaignOutboundCallConfigCmd.MarkFlagRequired("id")
	})
	connectcampaignsCmd.AddCommand(connectcampaigns_updateCampaignOutboundCallConfigCmd)
}
