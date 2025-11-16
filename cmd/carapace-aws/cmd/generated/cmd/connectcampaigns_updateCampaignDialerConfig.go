package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_updateCampaignDialerConfigCmd = &cobra.Command{
	Use:   "update-campaign-dialer-config",
	Short: "Updates the dialer config of a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_updateCampaignDialerConfigCmd).Standalone()

	connectcampaigns_updateCampaignDialerConfigCmd.Flags().String("dialer-config", "", "")
	connectcampaigns_updateCampaignDialerConfigCmd.Flags().String("id", "", "")
	connectcampaigns_updateCampaignDialerConfigCmd.MarkFlagRequired("dialer-config")
	connectcampaigns_updateCampaignDialerConfigCmd.MarkFlagRequired("id")
	connectcampaignsCmd.AddCommand(connectcampaigns_updateCampaignDialerConfigCmd)
}
