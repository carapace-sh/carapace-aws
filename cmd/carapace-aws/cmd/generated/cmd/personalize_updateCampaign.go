package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_updateCampaignCmd = &cobra.Command{
	Use:   "update-campaign",
	Short: "Updates a campaign to deploy a retrained solution version with an existing campaign, change your campaign's `minProvisionedTPS`, or modify your campaign's configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_updateCampaignCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_updateCampaignCmd).Standalone()

		personalize_updateCampaignCmd.Flags().String("campaign-arn", "", "The Amazon Resource Name (ARN) of the campaign.")
		personalize_updateCampaignCmd.Flags().String("campaign-config", "", "The configuration details of a campaign.")
		personalize_updateCampaignCmd.Flags().String("min-provisioned-tps", "", "Specifies the requested minimum provisioned transactions (recommendations) per second that Amazon Personalize will support.")
		personalize_updateCampaignCmd.Flags().String("solution-version-arn", "", "The Amazon Resource Name (ARN) of a new model to deploy.")
		personalize_updateCampaignCmd.MarkFlagRequired("campaign-arn")
	})
	personalizeCmd.AddCommand(personalize_updateCampaignCmd)
}
