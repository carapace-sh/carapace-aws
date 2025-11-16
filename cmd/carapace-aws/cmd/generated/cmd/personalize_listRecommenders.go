package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listRecommendersCmd = &cobra.Command{
	Use:   "list-recommenders",
	Short: "Returns a list of recommenders in a given Domain dataset group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listRecommendersCmd).Standalone()

	personalize_listRecommendersCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the Domain dataset group to list the recommenders for.")
	personalize_listRecommendersCmd.Flags().String("max-results", "", "The maximum number of recommenders to return.")
	personalize_listRecommendersCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListRecommenders` for getting the next set of recommenders (if they exist).")
	personalizeCmd.AddCommand(personalize_listRecommendersCmd)
}
