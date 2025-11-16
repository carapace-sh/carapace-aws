package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deletePartnerAppCmd = &cobra.Command{
	Use:   "delete-partner-app",
	Short: "Deletes a SageMaker Partner AI App.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deletePartnerAppCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deletePartnerAppCmd).Standalone()

		sagemaker_deletePartnerAppCmd.Flags().String("arn", "", "The ARN of the SageMaker Partner AI App to delete.")
		sagemaker_deletePartnerAppCmd.Flags().String("client-token", "", "A unique token that guarantees that the call to this API is idempotent.")
		sagemaker_deletePartnerAppCmd.MarkFlagRequired("arn")
	})
	sagemakerCmd.AddCommand(sagemaker_deletePartnerAppCmd)
}
