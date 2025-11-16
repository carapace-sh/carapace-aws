package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_describeCampaignCmd = &cobra.Command{
	Use:   "describe-campaign",
	Short: "Describes the given campaign, including its status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_describeCampaignCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_describeCampaignCmd).Standalone()

		personalize_describeCampaignCmd.Flags().String("campaign-arn", "", "The Amazon Resource Name (ARN) of the campaign.")
		personalize_describeCampaignCmd.MarkFlagRequired("campaign-arn")
	})
	personalizeCmd.AddCommand(personalize_describeCampaignCmd)
}
