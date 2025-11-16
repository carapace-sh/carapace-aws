package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_listFindingsCmd = &cobra.Command{
	Use:   "list-findings",
	Short: "Lists findings for your environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_listFindingsCmd).Standalone()

	inspector2_listFindingsCmd.Flags().String("filter-criteria", "", "Details on the filters to apply to your finding results.")
	inspector2_listFindingsCmd.Flags().String("max-results", "", "The maximum number of results the response can return.")
	inspector2_listFindingsCmd.Flags().String("next-token", "", "A token to use for paginating results that are returned in the response.")
	inspector2_listFindingsCmd.Flags().String("sort-criteria", "", "Details on the sort criteria to apply to your finding results.")
	inspector2Cmd.AddCommand(inspector2_listFindingsCmd)
}
