package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listDatasetsCmd = &cobra.Command{
	Use:   "list-datasets",
	Short: "Retrieves a paginated list of datasets for a specific target resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listDatasetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listDatasetsCmd).Standalone()

		iotsitewise_listDatasetsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listDatasetsCmd.Flags().String("next-token", "", "The token for the next set of results, or null if there are no additional results.")
		iotsitewise_listDatasetsCmd.Flags().String("source-type", "", "The type of data source for the dataset.")
		iotsitewise_listDatasetsCmd.MarkFlagRequired("source-type")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listDatasetsCmd)
}
