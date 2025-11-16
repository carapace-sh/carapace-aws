package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_startLendingAnalysisCmd = &cobra.Command{
	Use:   "start-lending-analysis",
	Short: "Starts the classification and analysis of an input document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_startLendingAnalysisCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(textract_startLendingAnalysisCmd).Standalone()

		textract_startLendingAnalysisCmd.Flags().String("client-request-token", "", "The idempotent token that you use to identify the start request.")
		textract_startLendingAnalysisCmd.Flags().String("document-location", "", "")
		textract_startLendingAnalysisCmd.Flags().String("job-tag", "", "An identifier that you specify to be included in the completion notification published to the Amazon SNS topic.")
		textract_startLendingAnalysisCmd.Flags().String("kmskey-id", "", "The KMS key used to encrypt the inference results.")
		textract_startLendingAnalysisCmd.Flags().String("notification-channel", "", "")
		textract_startLendingAnalysisCmd.Flags().String("output-config", "", "")
		textract_startLendingAnalysisCmd.MarkFlagRequired("document-location")
	})
	textractCmd.AddCommand(textract_startLendingAnalysisCmd)
}
