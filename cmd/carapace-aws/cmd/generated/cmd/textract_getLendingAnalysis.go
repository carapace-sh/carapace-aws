package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_getLendingAnalysisCmd = &cobra.Command{
	Use:   "get-lending-analysis",
	Short: "Gets the results for an Amazon Textract asynchronous operation that analyzes text in a lending document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_getLendingAnalysisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textract_getLendingAnalysisCmd).Standalone()

		textract_getLendingAnalysisCmd.Flags().String("job-id", "", "A unique identifier for the lending or text-detection job.")
		textract_getLendingAnalysisCmd.Flags().String("max-results", "", "The maximum number of results to return per paginated call.")
		textract_getLendingAnalysisCmd.Flags().String("next-token", "", "If the previous response was incomplete, Amazon Textract returns a pagination token in the response.")
		textract_getLendingAnalysisCmd.MarkFlagRequired("job-id")
	})
	textractCmd.AddCommand(textract_getLendingAnalysisCmd)
}
