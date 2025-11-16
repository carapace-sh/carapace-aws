package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listCampaignsCmd = &cobra.Command{
	Use:   "list-campaigns",
	Short: "Returns a list of campaigns that use the given solution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listCampaignsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_listCampaignsCmd).Standalone()

		personalize_listCampaignsCmd.Flags().String("max-results", "", "The maximum number of campaigns to return.")
		personalize_listCampaignsCmd.Flags().String("next-token", "", "A token returned from the previous call to [ListCampaigns](https://docs.aws.amazon.com/personalize/latest/dg/API_ListCampaigns.html) for getting the next set of campaigns (if they exist).")
		personalize_listCampaignsCmd.Flags().String("solution-arn", "", "The Amazon Resource Name (ARN) of the solution to list the campaigns for.")
	})
	personalizeCmd.AddCommand(personalize_listCampaignsCmd)
}
