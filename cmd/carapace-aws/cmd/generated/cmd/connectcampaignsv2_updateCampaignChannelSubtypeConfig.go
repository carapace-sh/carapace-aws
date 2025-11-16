package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_updateCampaignChannelSubtypeConfigCmd = &cobra.Command{
	Use:   "update-campaign-channel-subtype-config",
	Short: "Updates the channel subtype config of a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_updateCampaignChannelSubtypeConfigCmd).Standalone()

	connectcampaignsv2_updateCampaignChannelSubtypeConfigCmd.Flags().String("channel-subtype-config", "", "")
	connectcampaignsv2_updateCampaignChannelSubtypeConfigCmd.Flags().String("id", "", "")
	connectcampaignsv2_updateCampaignChannelSubtypeConfigCmd.MarkFlagRequired("channel-subtype-config")
	connectcampaignsv2_updateCampaignChannelSubtypeConfigCmd.MarkFlagRequired("id")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_updateCampaignChannelSubtypeConfigCmd)
}
