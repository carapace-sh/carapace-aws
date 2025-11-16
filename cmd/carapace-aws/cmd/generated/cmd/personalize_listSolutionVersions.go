package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listSolutionVersionsCmd = &cobra.Command{
	Use:   "list-solution-versions",
	Short: "Returns a list of solution versions for the given solution.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listSolutionVersionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(personalize_listSolutionVersionsCmd).Standalone()

		personalize_listSolutionVersionsCmd.Flags().String("max-results", "", "The maximum number of solution versions to return.")
		personalize_listSolutionVersionsCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListSolutionVersions` for getting the next set of solution versions (if they exist).")
		personalize_listSolutionVersionsCmd.Flags().String("solution-arn", "", "The Amazon Resource Name (ARN) of the solution.")
	})
	personalizeCmd.AddCommand(personalize_listSolutionVersionsCmd)
}
