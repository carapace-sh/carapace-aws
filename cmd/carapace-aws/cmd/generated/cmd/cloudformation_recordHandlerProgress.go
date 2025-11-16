package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_recordHandlerProgressCmd = &cobra.Command{
	Use:   "record-handler-progress",
	Short: "Reports progress of a resource handler to CloudFormation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_recordHandlerProgressCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_recordHandlerProgressCmd).Standalone()

		cloudformation_recordHandlerProgressCmd.Flags().String("bearer-token", "", "Reserved for use by the [CloudFormation CLI](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).")
		cloudformation_recordHandlerProgressCmd.Flags().String("client-request-token", "", "Reserved for use by the [CloudFormation CLI](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).")
		cloudformation_recordHandlerProgressCmd.Flags().String("current-operation-status", "", "Reserved for use by the [CloudFormation CLI](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).")
		cloudformation_recordHandlerProgressCmd.Flags().String("error-code", "", "Reserved for use by the [CloudFormation CLI](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).")
		cloudformation_recordHandlerProgressCmd.Flags().String("operation-status", "", "Reserved for use by the [CloudFormation CLI](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).")
		cloudformation_recordHandlerProgressCmd.Flags().String("resource-model", "", "Reserved for use by the [CloudFormation CLI](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).")
		cloudformation_recordHandlerProgressCmd.Flags().String("status-message", "", "Reserved for use by the [CloudFormation CLI](https://docs.aws.amazon.com/cloudformation-cli/latest/userguide/what-is-cloudformation-cli.html).")
		cloudformation_recordHandlerProgressCmd.MarkFlagRequired("bearer-token")
		cloudformation_recordHandlerProgressCmd.MarkFlagRequired("operation-status")
	})
	cloudformationCmd.AddCommand(cloudformation_recordHandlerProgressCmd)
}
