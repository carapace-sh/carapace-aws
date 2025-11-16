package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_getLendingAnalysisSummaryCmd = &cobra.Command{
	Use:   "get-lending-analysis-summary",
	Short: "Gets summarized results for the `StartLendingAnalysis` operation, which analyzes text in a lending document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_getLendingAnalysisSummaryCmd).Standalone()

	textract_getLendingAnalysisSummaryCmd.Flags().String("job-id", "", "A unique identifier for the lending or text-detection job.")
	textract_getLendingAnalysisSummaryCmd.MarkFlagRequired("job-id")
	textractCmd.AddCommand(textract_getLendingAnalysisSummaryCmd)
}
