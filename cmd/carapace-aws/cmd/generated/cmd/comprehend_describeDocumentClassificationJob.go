package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_describeDocumentClassificationJobCmd = &cobra.Command{
	Use:   "describe-document-classification-job",
	Short: "Gets the properties associated with a document classification job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_describeDocumentClassificationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehend_describeDocumentClassificationJobCmd).Standalone()

		comprehend_describeDocumentClassificationJobCmd.Flags().String("job-id", "", "The identifier that Amazon Comprehend generated for the job.")
		comprehend_describeDocumentClassificationJobCmd.MarkFlagRequired("job-id")
	})
	comprehendCmd.AddCommand(comprehend_describeDocumentClassificationJobCmd)
}
