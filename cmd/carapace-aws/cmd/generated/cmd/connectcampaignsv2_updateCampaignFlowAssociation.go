package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_updateCampaignFlowAssociationCmd = &cobra.Command{
	Use:   "update-campaign-flow-association",
	Short: "Updates the campaign flow associated with a campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_updateCampaignFlowAssociationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_updateCampaignFlowAssociationCmd).Standalone()

		connectcampaignsv2_updateCampaignFlowAssociationCmd.Flags().String("connect-campaign-flow-arn", "", "")
		connectcampaignsv2_updateCampaignFlowAssociationCmd.Flags().String("id", "", "")
		connectcampaignsv2_updateCampaignFlowAssociationCmd.MarkFlagRequired("connect-campaign-flow-arn")
		connectcampaignsv2_updateCampaignFlowAssociationCmd.MarkFlagRequired("id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_updateCampaignFlowAssociationCmd)
}
