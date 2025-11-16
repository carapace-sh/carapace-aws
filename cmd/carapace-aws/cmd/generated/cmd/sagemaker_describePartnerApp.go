package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describePartnerAppCmd = &cobra.Command{
	Use:   "describe-partner-app",
	Short: "Gets information about a SageMaker Partner AI App.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describePartnerAppCmd).Standalone()

	sagemaker_describePartnerAppCmd.Flags().String("arn", "", "The ARN of the SageMaker Partner AI App to describe.")
	sagemaker_describePartnerAppCmd.Flags().Bool("include-available-upgrade", false, "When set to `TRUE`, the response includes available upgrade information for the SageMaker Partner AI App.")
	sagemaker_describePartnerAppCmd.Flags().Bool("no-include-available-upgrade", false, "When set to `TRUE`, the response includes available upgrade information for the SageMaker Partner AI App.")
	sagemaker_describePartnerAppCmd.MarkFlagRequired("arn")
	sagemaker_describePartnerAppCmd.Flag("no-include-available-upgrade").Hidden = true
	sagemakerCmd.AddCommand(sagemaker_describePartnerAppCmd)
}
