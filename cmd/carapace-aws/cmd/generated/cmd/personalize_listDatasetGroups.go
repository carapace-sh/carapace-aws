package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var personalize_listDatasetGroupsCmd = &cobra.Command{
	Use:   "list-dataset-groups",
	Short: "Returns a list of dataset groups.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(personalize_listDatasetGroupsCmd).Standalone()

	personalize_listDatasetGroupsCmd.Flags().String("max-results", "", "The maximum number of dataset groups to return.")
	personalize_listDatasetGroupsCmd.Flags().String("next-token", "", "A token returned from the previous call to `ListDatasetGroups` for getting the next set of dataset groups (if they exist).")
	personalizeCmd.AddCommand(personalize_listDatasetGroupsCmd)
}
