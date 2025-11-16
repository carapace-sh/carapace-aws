package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_createCampaignCmd = &cobra.Command{
	Use:   "create-campaign",
	Short: "You incur campaign costs while it is active.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_createCampaignCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_createCampaignCmd).Standalone()

		personalize_createCampaignCmd.Flags().String("campaign-config", "", "The configuration details of a campaign.")
		personalize_createCampaignCmd.Flags().String("min-provisioned-tps", "", "Specifies the requested minimum provisioned transactions (recommendations) per second that Amazon Personalize will support.")
		personalize_createCampaignCmd.Flags().String("name", "", "A name for the new campaign.")
		personalize_createCampaignCmd.Flags().String("solution-version-arn", "", "The Amazon Resource Name (ARN) of the trained model to deploy with the campaign.")
		personalize_createCampaignCmd.Flags().String("tags", "", "A list of [tags](https://docs.aws.amazon.com/personalize/latest/dg/tagging-resources.html) to apply to the campaign.")
		personalize_createCampaignCmd.MarkFlagRequired("name")
		personalize_createCampaignCmd.MarkFlagRequired("solution-version-arn")
	})
	personalizeCmd.AddCommand(personalize_createCampaignCmd)
}
