package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listDatasetsCmd = &cobra.Command{
	Use:   "list-datasets",
	Short: "List the datasets that you have configured in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listDatasetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_listDatasetsCmd).Standalone()

		comprehend_listDatasetsCmd.Flags().String("filter", "", "Filters the datasets to be returned in the response.")
		comprehend_listDatasetsCmd.Flags().String("flywheel-arn", "", "The Amazon Resource Number (ARN) of the flywheel.")
		comprehend_listDatasetsCmd.Flags().String("max-results", "", "Maximum number of results to return in a response.")
		comprehend_listDatasetsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	})
	comprehendCmd.AddCommand(comprehend_listDatasetsCmd)
}
