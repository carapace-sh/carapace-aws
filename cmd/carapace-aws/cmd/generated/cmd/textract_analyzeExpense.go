package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_analyzeExpenseCmd = &cobra.Command{
	Use:   "analyze-expense",
	Short: "`AnalyzeExpense` synchronously analyzes an input document for financially related relationships between text.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_analyzeExpenseCmd).Standalone()

	textract_analyzeExpenseCmd.Flags().String("document", "", "")
	textract_analyzeExpenseCmd.MarkFlagRequired("document")
	textractCmd.AddCommand(textract_analyzeExpenseCmd)
}
