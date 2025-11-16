package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_listRunGroupsCmd = &cobra.Command{
	Use:   "list-run-groups",
	Short: "Retrieves a list of all run groups and returns the metadata for each run group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_listRunGroupsCmd).Standalone()

	omics_listRunGroupsCmd.Flags().String("max-results", "", "The maximum number of run groups to return in one page of results.")
	omics_listRunGroupsCmd.Flags().String("name", "", "The run groups' name.")
	omics_listRunGroupsCmd.Flags().String("starting-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
	omicsCmd.AddCommand(omics_listRunGroupsCmd)
}
