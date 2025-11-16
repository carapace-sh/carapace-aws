package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listSolutionsCmd = &cobra.Command{
	Use:   "list-solutions",
	Short: "Returns a list of solutions in a given dataset group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listSolutionsCmd).Standalone()

	personalize_listSolutionsCmd.Flags().String("dataset-group-arn", "", "The Amazon Resource Name (ARN) of the dataset group.")
	personalize_listSolutionsCmd.Flags().String("max-results", "", "The maximum number of solutions to return.")
	personalize_listSolutionsCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListSolutions` for getting the next set of solutions (if they exist).")
	personalizeCmd.AddCommand(personalize_listSolutionsCmd)
}
