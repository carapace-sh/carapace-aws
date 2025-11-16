package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listEntityRecognizerSummariesCmd = &cobra.Command{
	Use:   "list-entity-recognizer-summaries",
	Short: "Gets a list of summaries for the entity recognizers that you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listEntityRecognizerSummariesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_listEntityRecognizerSummariesCmd).Standalone()

		comprehend_listEntityRecognizerSummariesCmd.Flags().String("max-results", "", "The maximum number of results to return on each page.")
		comprehend_listEntityRecognizerSummariesCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	})
	comprehendCmd.AddCommand(comprehend_listEntityRecognizerSummariesCmd)
}
