package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_createCampaignCmd = &cobra.Command{
	Use:   "create-campaign",
	Short: "Creates a campaign for the specified Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_createCampaignCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaigns_createCampaignCmd).Standalone()

		connectcampaigns_createCampaignCmd.Flags().String("connect-instance-id", "", "")
		connectcampaigns_createCampaignCmd.Flags().String("dialer-config", "", "")
		connectcampaigns_createCampaignCmd.Flags().String("name", "", "")
		connectcampaigns_createCampaignCmd.Flags().String("outbound-call-config", "", "")
		connectcampaigns_createCampaignCmd.Flags().String("tags", "", "")
		connectcampaigns_createCampaignCmd.MarkFlagRequired("connect-instance-id")
		connectcampaigns_createCampaignCmd.MarkFlagRequired("dialer-config")
		connectcampaigns_createCampaignCmd.MarkFlagRequired("name")
		connectcampaigns_createCampaignCmd.MarkFlagRequired("outbound-call-config")
	})
	connectcampaignsCmd.AddCommand(connectcampaigns_createCampaignCmd)
}
