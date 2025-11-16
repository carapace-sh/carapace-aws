package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_createChangeSetCmd = &cobra.Command{
	Use:   "create-change-set",
	Short: "Creates a list of changes that will be applied to a stack so that you can review the changes before executing them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_createChangeSetCmd).Standalone()

	cloudformation_createChangeSetCmd.Flags().String("capabilities", "", "In some cases, you must explicitly acknowledge that your stack template contains certain capabilities in order for CloudFormation to create the stack.")
	cloudformation_createChangeSetCmd.Flags().String("change-set-name", "", "The name of the change set.")
	cloudformation_createChangeSetCmd.Flags().String("change-set-type", "", "The type of change set operation.")
	cloudformation_createChangeSetCmd.Flags().String("client-token", "", "A unique identifier for this `CreateChangeSet` request.")
	cloudformation_createChangeSetCmd.Flags().String("description", "", "A description to help you identify this change set.")
	cloudformation_createChangeSetCmd.Flags().String("import-existing-resources", "", "Indicates if the change set auto-imports resources that already exist.")
	cloudformation_createChangeSetCmd.Flags().String("include-nested-stacks", "", "Creates a change set for the all nested stacks specified in the template.")
	cloudformation_createChangeSetCmd.Flags().String("notification-arns", "", "The Amazon Resource Names (ARNs) of Amazon SNS topics that CloudFormation associates with the stack.")
	cloudformation_createChangeSetCmd.Flags().String("on-stack-failure", "", "Determines what action will be taken if stack creation fails.")
	cloudformation_createChangeSetCmd.Flags().String("parameters", "", "A list of `Parameter` structures that specify input parameters for the change set.")
	cloudformation_createChangeSetCmd.Flags().String("resource-types", "", "Specifies which resource types you can work with, such as `AWS::EC2::Instance` or `Custom::MyCustomInstance`.")
	cloudformation_createChangeSetCmd.Flags().String("resources-to-import", "", "The resources to import into your stack.")
	cloudformation_createChangeSetCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that CloudFormation assumes when executing the change set.")
	cloudformation_createChangeSetCmd.Flags().String("rollback-configuration", "", "The rollback triggers for CloudFormation to monitor during stack creation and updating operations, and for the specified monitoring period afterwards.")
	cloudformation_createChangeSetCmd.Flags().String("stack-name", "", "The name or the unique ID of the stack for which you are creating a change set.")
	cloudformation_createChangeSetCmd.Flags().String("tags", "", "Key-value pairs to associate with this stack.")
	cloudformation_createChangeSetCmd.Flags().String("template-body", "", "A structure that contains the body of the revised template, with a minimum length of 1 byte and a maximum length of 51,200 bytes.")
	cloudformation_createChangeSetCmd.Flags().String("template-url", "", "The URL of the file that contains the revised template.")
	cloudformation_createChangeSetCmd.Flags().String("use-previous-template", "", "Whether to reuse the template that's associated with the stack to create the change set.")
	cloudformation_createChangeSetCmd.MarkFlagRequired("change-set-name")
	cloudformation_createChangeSetCmd.MarkFlagRequired("stack-name")
	cloudformationCmd.AddCommand(cloudformation_createChangeSetCmd)
}
