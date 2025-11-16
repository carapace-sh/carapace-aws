package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_listDatasetsCmd = &cobra.Command{
	Use:   "list-datasets",
	Short: "Lists all of the DataBrew datasets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_listDatasetsCmd).Standalone()

	databrew_listDatasetsCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	databrew_listDatasetsCmd.Flags().String("next-token", "", "The token returned by a previous call to retrieve the next set of results.")
	databrewCmd.AddCommand(databrew_listDatasetsCmd)
}
