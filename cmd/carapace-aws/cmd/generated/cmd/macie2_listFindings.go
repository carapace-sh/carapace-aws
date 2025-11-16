package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_listFindingsCmd = &cobra.Command{
	Use:   "list-findings",
	Short: "Retrieves a subset of information about one or more findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_listFindingsCmd).Standalone()

	macie2_listFindingsCmd.Flags().String("finding-criteria", "", "The criteria to use to filter the results.")
	macie2_listFindingsCmd.Flags().String("max-results", "", "The maximum number of items to include in each page of the response.")
	macie2_listFindingsCmd.Flags().String("next-token", "", "The nextToken string that specifies which page of results to return in a paginated response.")
	macie2_listFindingsCmd.Flags().String("sort-criteria", "", "The criteria to use to sort the results.")
	macie2Cmd.AddCommand(macie2_listFindingsCmd)
}
