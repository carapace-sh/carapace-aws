package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_deleteCampaignCmd = &cobra.Command{
	Use:   "delete-campaign",
	Short: "Removes a campaign by deleting the solution deployment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_deleteCampaignCmd).Standalone()

	personalize_deleteCampaignCmd.Flags().String("campaign-arn", "", "The Amazon Resource Name (ARN) of the campaign to delete.")
	personalize_deleteCampaignCmd.MarkFlagRequired("campaign-arn")
	personalizeCmd.AddCommand(personalize_deleteCampaignCmd)
}
