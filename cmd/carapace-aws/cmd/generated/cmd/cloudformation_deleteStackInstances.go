package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_deleteStackInstancesCmd = &cobra.Command{
	Use:   "delete-stack-instances",
	Short: "Deletes stack instances for the specified accounts, in the specified Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_deleteStackInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_deleteStackInstancesCmd).Standalone()

		cloudformation_deleteStackInstancesCmd.Flags().String("accounts", "", "\\[Self-managed permissions] The account IDs of the Amazon Web Services accounts that you want to delete stack instances for.")
		cloudformation_deleteStackInstancesCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
		cloudformation_deleteStackInstancesCmd.Flags().String("deployment-targets", "", "\\[Service-managed permissions] The Organizations accounts from which to delete stack instances.")
		cloudformation_deleteStackInstancesCmd.Flags().String("operation-id", "", "The unique identifier for this StackSet operation.")
		cloudformation_deleteStackInstancesCmd.Flags().String("operation-preferences", "", "Preferences for how CloudFormation performs this StackSet operation.")
		cloudformation_deleteStackInstancesCmd.Flags().String("regions", "", "The Amazon Web Services Regions where you want to delete StackSet instances.")
		cloudformation_deleteStackInstancesCmd.Flags().String("retain-stacks", "", "Removes the stack instances from the specified StackSet, but doesn't delete the stacks.")
		cloudformation_deleteStackInstancesCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet that you want to delete stack instances for.")
		cloudformation_deleteStackInstancesCmd.MarkFlagRequired("regions")
		cloudformation_deleteStackInstancesCmd.MarkFlagRequired("retain-stacks")
		cloudformation_deleteStackInstancesCmd.MarkFlagRequired("stack-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_deleteStackInstancesCmd)
}
