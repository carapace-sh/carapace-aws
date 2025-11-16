package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_updateStackCmd = &cobra.Command{
	Use:   "update-stack",
	Short: "Updates a stack as specified in the template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_updateStackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_updateStackCmd).Standalone()

		cloudformation_updateStackCmd.Flags().String("capabilities", "", "In some cases, you must explicitly acknowledge that your stack template contains certain capabilities in order for CloudFormation to update the stack.")
		cloudformation_updateStackCmd.Flags().String("client-request-token", "", "A unique identifier for this `UpdateStack` request.")
		cloudformation_updateStackCmd.Flags().String("disable-rollback", "", "Preserve the state of previously provisioned resources when an operation fails.")
		cloudformation_updateStackCmd.Flags().String("notification-arns", "", "Amazon Simple Notification Service topic Amazon Resource Names (ARNs) that CloudFormation associates with the stack.")
		cloudformation_updateStackCmd.Flags().String("parameters", "", "A list of `Parameter` structures that specify input parameters for the stack.")
		cloudformation_updateStackCmd.Flags().String("resource-types", "", "Specifies which resource types you can work with, such as `AWS::EC2::Instance` or `Custom::MyCustomInstance`.")
		cloudformation_updateStackCmd.Flags().String("retain-except-on-create", "", "When set to `true`, newly created resources are deleted when the operation rolls back.")
		cloudformation_updateStackCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that CloudFormation assumes to update the stack.")
		cloudformation_updateStackCmd.Flags().String("rollback-configuration", "", "The rollback triggers for CloudFormation to monitor during stack creation and updating operations, and for the specified monitoring period afterwards.")
		cloudformation_updateStackCmd.Flags().String("stack-name", "", "The name or unique stack ID of the stack to update.")
		cloudformation_updateStackCmd.Flags().String("stack-policy-body", "", "Structure that contains a new stack policy body.")
		cloudformation_updateStackCmd.Flags().String("stack-policy-during-update-body", "", "Structure that contains the temporary overriding stack policy body.")
		cloudformation_updateStackCmd.Flags().String("stack-policy-during-update-url", "", "Location of a file that contains the temporary overriding stack policy.")
		cloudformation_updateStackCmd.Flags().String("stack-policy-url", "", "Location of a file that contains the updated stack policy.")
		cloudformation_updateStackCmd.Flags().String("tags", "", "Key-value pairs to associate with this stack.")
		cloudformation_updateStackCmd.Flags().String("template-body", "", "Structure that contains the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.")
		cloudformation_updateStackCmd.Flags().String("template-url", "", "The URL of a file that contains the template body.")
		cloudformation_updateStackCmd.Flags().String("use-previous-template", "", "Reuse the existing template that is associated with the stack that you are updating.")
		cloudformation_updateStackCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_updateStackCmd)
}
