package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_createStackCmd = &cobra.Command{
	Use:   "create-stack",
	Short: "Creates a stack as specified in the template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_createStackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_createStackCmd).Standalone()

		cloudformation_createStackCmd.Flags().String("capabilities", "", "In some cases, you must explicitly acknowledge that your stack template contains certain capabilities in order for CloudFormation to create the stack.")
		cloudformation_createStackCmd.Flags().String("client-request-token", "", "A unique identifier for this `CreateStack` request.")
		cloudformation_createStackCmd.Flags().String("disable-rollback", "", "Set to `true` to disable rollback of the stack if stack creation failed.")
		cloudformation_createStackCmd.Flags().String("enable-termination-protection", "", "Whether to enable termination protection on the specified stack.")
		cloudformation_createStackCmd.Flags().String("notification-arns", "", "The Amazon SNS topic ARNs to publish stack related events.")
		cloudformation_createStackCmd.Flags().String("on-failure", "", "Determines what action will be taken if stack creation fails.")
		cloudformation_createStackCmd.Flags().String("parameters", "", "A list of `Parameter` structures that specify input parameters for the stack.")
		cloudformation_createStackCmd.Flags().String("resource-types", "", "Specifies which resource types you can work with, such as `AWS::EC2::Instance` or `Custom::MyCustomInstance`.")
		cloudformation_createStackCmd.Flags().String("retain-except-on-create", "", "When set to `true`, newly created resources are deleted when the operation rolls back.")
		cloudformation_createStackCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that CloudFormation assumes to create the stack.")
		cloudformation_createStackCmd.Flags().String("rollback-configuration", "", "The rollback triggers for CloudFormation to monitor during stack creation and updating operations, and for the specified monitoring period afterwards.")
		cloudformation_createStackCmd.Flags().String("stack-name", "", "The name that's associated with the stack.")
		cloudformation_createStackCmd.Flags().String("stack-policy-body", "", "Structure that contains the stack policy body.")
		cloudformation_createStackCmd.Flags().String("stack-policy-url", "", "Location of a file that contains the stack policy.")
		cloudformation_createStackCmd.Flags().String("tags", "", "Key-value pairs to associate with this stack.")
		cloudformation_createStackCmd.Flags().String("template-body", "", "Structure that contains the template body with a minimum length of 1 byte and a maximum length of 51,200 bytes.")
		cloudformation_createStackCmd.Flags().String("template-url", "", "The URL of a file that contains the template body.")
		cloudformation_createStackCmd.Flags().String("timeout-in-minutes", "", "The amount of time that can pass before the stack status becomes `CREATE_FAILED`; if `DisableRollback` is not set or is set to `false`, the stack will be rolled back.")
		cloudformation_createStackCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_createStackCmd)
}
