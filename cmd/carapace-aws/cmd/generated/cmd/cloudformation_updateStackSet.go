package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_updateStackSetCmd = &cobra.Command{
	Use:   "update-stack-set",
	Short: "Updates the StackSet and associated stack instances in the specified accounts and Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_updateStackSetCmd).Standalone()

	cloudformation_updateStackSetCmd.Flags().String("accounts", "", "\\[Self-managed permissions] The accounts in which to update associated stack instances.")
	cloudformation_updateStackSetCmd.Flags().String("administration-role-arn", "", "\\[Self-managed permissions] The Amazon Resource Name (ARN) of the IAM role to use to update this StackSet.")
	cloudformation_updateStackSetCmd.Flags().String("auto-deployment", "", "\\[Service-managed permissions] Describes whether StackSets automatically deploys to Organizations accounts that are added to a target organization or organizational unit (OU).")
	cloudformation_updateStackSetCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
	cloudformation_updateStackSetCmd.Flags().String("capabilities", "", "In some cases, you must explicitly acknowledge that your stack template contains certain capabilities in order for CloudFormation to update the StackSet and its associated stack instances.")
	cloudformation_updateStackSetCmd.Flags().String("deployment-targets", "", "\\[Service-managed permissions] The Organizations accounts in which to update associated stack instances.")
	cloudformation_updateStackSetCmd.Flags().String("description", "", "A brief description of updates that you are making.")
	cloudformation_updateStackSetCmd.Flags().String("execution-role-name", "", "\\[Self-managed permissions] The name of the IAM execution role to use to update the stack set.")
	cloudformation_updateStackSetCmd.Flags().String("managed-execution", "", "Describes whether CloudFormation performs non-conflicting operations concurrently and queues conflicting operations.")
	cloudformation_updateStackSetCmd.Flags().String("operation-id", "", "The unique ID for this StackSet operation.")
	cloudformation_updateStackSetCmd.Flags().String("operation-preferences", "", "Preferences for how CloudFormation performs this StackSet operation.")
	cloudformation_updateStackSetCmd.Flags().String("parameters", "", "A list of input parameters for the StackSet template.")
	cloudformation_updateStackSetCmd.Flags().String("permission-model", "", "Describes how the IAM roles required for StackSet operations are created.")
	cloudformation_updateStackSetCmd.Flags().String("regions", "", "The Amazon Web Services Regions in which to update associated stack instances.")
	cloudformation_updateStackSetCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet that you want to update.")
	cloudformation_updateStackSetCmd.Flags().String("tags", "", "The key-value pairs to associate with this StackSet and the stacks created from it.")
	cloudformation_updateStackSetCmd.Flags().String("template-body", "", "The structure that contains the template body, with a minimum length of 1 byte and a maximum length of 51,200 bytes.")
	cloudformation_updateStackSetCmd.Flags().String("template-url", "", "The URL of a file that contains the template body.")
	cloudformation_updateStackSetCmd.Flags().String("use-previous-template", "", "Use the existing template that's associated with the StackSet that you're updating.")
	cloudformation_updateStackSetCmd.MarkFlagRequired("stack-set-name")
	cloudformationCmd.AddCommand(cloudformation_updateStackSetCmd)
}
