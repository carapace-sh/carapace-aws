package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var textract_startExpenseAnalysisCmd = &cobra.Command{
	Use:   "start-expense-analysis",
	Short: "Starts the asynchronous analysis of invoices or receipts for data like contact information, items purchased, and vendor names.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(textract_startExpenseAnalysisCmd).Standalone()

	textract_startExpenseAnalysisCmd.Flags().String("client-request-token", "", "The idempotent token that's used to identify the start request.")
	textract_startExpenseAnalysisCmd.Flags().String("document-location", "", "The location of the document to be processed.")
	textract_startExpenseAnalysisCmd.Flags().String("job-tag", "", "An identifier you specify that's included in the completion notification published to the Amazon SNS topic.")
	textract_startExpenseAnalysisCmd.Flags().String("kmskey-id", "", "The KMS key used to encrypt the inference results.")
	textract_startExpenseAnalysisCmd.Flags().String("notification-channel", "", "The Amazon SNS topic ARN that you want Amazon Textract to publish the completion status of the operation to.")
	textract_startExpenseAnalysisCmd.Flags().String("output-config", "", "Sets if the output will go to a customer defined bucket.")
	textract_startExpenseAnalysisCmd.MarkFlagRequired("document-location")
	textractCmd.AddCommand(textract_startExpenseAnalysisCmd)
}
