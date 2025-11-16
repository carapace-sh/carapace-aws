package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_deleteCampaignChannelSubtypeConfigCmd = &cobra.Command{
	Use:   "delete-campaign-channel-subtype-config",
	Short: "Deletes the channel subtype config of a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_deleteCampaignChannelSubtypeConfigCmd).Standalone()

	connectcampaignsv2_deleteCampaignChannelSubtypeConfigCmd.Flags().String("channel-subtype", "", "")
	connectcampaignsv2_deleteCampaignChannelSubtypeConfigCmd.Flags().String("id", "", "")
	connectcampaignsv2_deleteCampaignChannelSubtypeConfigCmd.MarkFlagRequired("channel-subtype")
	connectcampaignsv2_deleteCampaignChannelSubtypeConfigCmd.MarkFlagRequired("id")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_deleteCampaignChannelSubtypeConfigCmd)
}
