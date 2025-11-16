package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_getExpenseAnalysisCmd = &cobra.Command{
	Use:   "get-expense-analysis",
	Short: "Gets the results for an Amazon Textract asynchronous operation that analyzes invoices and receipts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_getExpenseAnalysisCmd).Standalone()

	textract_getExpenseAnalysisCmd.Flags().String("job-id", "", "A unique identifier for the text detection job.")
	textract_getExpenseAnalysisCmd.Flags().String("max-results", "", "The maximum number of results to return per paginated call.")
	textract_getExpenseAnalysisCmd.Flags().String("next-token", "", "If the previous response was incomplete (because there are more blocks to retrieve), Amazon Textract returns a pagination token in the response.")
	textract_getExpenseAnalysisCmd.MarkFlagRequired("job-id")
	textractCmd.AddCommand(textract_getExpenseAnalysisCmd)
}
