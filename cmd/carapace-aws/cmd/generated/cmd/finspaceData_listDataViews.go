package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_listDataViewsCmd = &cobra.Command{
	Use:   "list-data-views",
	Short: "Lists all available Dataviews for a Dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_listDataViewsCmd).Standalone()

	finspaceData_listDataViewsCmd.Flags().String("dataset-id", "", "The unique identifier of the Dataset for which to retrieve Dataviews.")
	finspaceData_listDataViewsCmd.Flags().String("max-results", "", "The maximum number of results per page.")
	finspaceData_listDataViewsCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspaceData_listDataViewsCmd.MarkFlagRequired("dataset-id")
	finspaceDataCmd.AddCommand(finspaceData_listDataViewsCmd)
}
