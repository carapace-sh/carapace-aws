package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_updateStackInstancesCmd = &cobra.Command{
	Use:   "update-stack-instances",
	Short: "Updates the parameter values for stack instances for the specified accounts, within the specified Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_updateStackInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_updateStackInstancesCmd).Standalone()

		cloudformation_updateStackInstancesCmd.Flags().String("accounts", "", "\\[Self-managed permissions] The account IDs of one or more Amazon Web Services accounts in which you want to update parameter values for stack instances.")
		cloudformation_updateStackInstancesCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
		cloudformation_updateStackInstancesCmd.Flags().String("deployment-targets", "", "\\[Service-managed permissions] The Organizations accounts in which you want to update parameter values for stack instances.")
		cloudformation_updateStackInstancesCmd.Flags().String("operation-id", "", "The unique identifier for this StackSet operation.")
		cloudformation_updateStackInstancesCmd.Flags().String("operation-preferences", "", "Preferences for how CloudFormation performs this StackSet operation.")
		cloudformation_updateStackInstancesCmd.Flags().String("parameter-overrides", "", "A list of input parameters whose values you want to update for the specified stack instances.")
		cloudformation_updateStackInstancesCmd.Flags().String("regions", "", "The names of one or more Amazon Web Services Regions in which you want to update parameter values for stack instances.")
		cloudformation_updateStackInstancesCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet associated with the stack instances.")
		cloudformation_updateStackInstancesCmd.MarkFlagRequired("regions")
		cloudformation_updateStackInstancesCmd.MarkFlagRequired("stack-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_updateStackInstancesCmd)
}
