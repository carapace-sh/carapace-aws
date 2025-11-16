package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_createStackSetCmd = &cobra.Command{
	Use:   "create-stack-set",
	Short: "Creates a StackSet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_createStackSetCmd).Standalone()

	cloudformation_createStackSetCmd.Flags().String("administration-role-arn", "", "The Amazon Resource Name (ARN) of the IAM role to use to create this StackSet.")
	cloudformation_createStackSetCmd.Flags().String("auto-deployment", "", "Describes whether StackSets automatically deploys to Organizations accounts that are added to the target organization or organizational unit (OU).")
	cloudformation_createStackSetCmd.Flags().String("call-as", "", "Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
	cloudformation_createStackSetCmd.Flags().String("capabilities", "", "In some cases, you must explicitly acknowledge that your StackSet template contains certain capabilities in order for CloudFormation to create the StackSet and related stack instances.")
	cloudformation_createStackSetCmd.Flags().String("client-request-token", "", "A unique identifier for this `CreateStackSet` request.")
	cloudformation_createStackSetCmd.Flags().String("description", "", "A description of the StackSet.")
	cloudformation_createStackSetCmd.Flags().String("execution-role-name", "", "The name of the IAM execution role to use to create the StackSet.")
	cloudformation_createStackSetCmd.Flags().String("managed-execution", "", "Describes whether CloudFormation performs non-conflicting operations concurrently and queues conflicting operations.")
	cloudformation_createStackSetCmd.Flags().String("parameters", "", "The input parameters for the StackSet template.")
	cloudformation_createStackSetCmd.Flags().String("permission-model", "", "Describes how the IAM roles required for StackSet operations are created.")
	cloudformation_createStackSetCmd.Flags().String("stack-id", "", "The stack ID you are importing into a new StackSet.")
	cloudformation_createStackSetCmd.Flags().String("stack-set-name", "", "The name to associate with the StackSet.")
	cloudformation_createStackSetCmd.Flags().String("tags", "", "The key-value pairs to associate with this StackSet and the stacks created from it.")
	cloudformation_createStackSetCmd.Flags().String("template-body", "", "The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.")
	cloudformation_createStackSetCmd.Flags().String("template-url", "", "The URL of a file that contains the template body.")
	cloudformation_createStackSetCmd.MarkFlagRequired("stack-set-name")
	cloudformationCmd.AddCommand(cloudformation_createStackSetCmd)
}
