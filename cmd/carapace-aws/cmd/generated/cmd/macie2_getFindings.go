package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getFindingsCmd = &cobra.Command{
	Use:   "get-findings",
	Short: "Retrieves the details of one or more findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getFindingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_getFindingsCmd).Standalone()

		macie2_getFindingsCmd.Flags().String("finding-ids", "", "An array of strings that lists the unique identifiers for the findings to retrieve.")
		macie2_getFindingsCmd.Flags().String("sort-criteria", "", "The criteria for sorting the results of the request.")
		macie2_getFindingsCmd.MarkFlagRequired("finding-ids")
	})
	macie2Cmd.AddCommand(macie2_getFindingsCmd)
}
