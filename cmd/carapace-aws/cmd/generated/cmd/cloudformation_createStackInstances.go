package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_createStackInstancesCmd = &cobra.Command{
	Use:   "create-stack-instances",
	Short: "Creates stack instances for the specified accounts, within the specified Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_createStackInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_createStackInstancesCmd).Standalone()

		cloudformation_createStackInstancesCmd.Flags().String("accounts", "", "\\[Self-managed permissions] The account IDs of one or more Amazon Web Services accounts that you want to create stack instances in the specified Region(s) for.")
		cloudformation_createStackInstancesCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
		cloudformation_createStackInstancesCmd.Flags().String("deployment-targets", "", "\\[Service-managed permissions] The Organizations accounts in which to create stack instances in the specified Amazon Web Services Regions.")
		cloudformation_createStackInstancesCmd.Flags().String("operation-id", "", "The unique identifier for this StackSet operation.")
		cloudformation_createStackInstancesCmd.Flags().String("operation-preferences", "", "Preferences for how CloudFormation performs this StackSet operation.")
		cloudformation_createStackInstancesCmd.Flags().String("parameter-overrides", "", "A list of StackSet parameters whose values you want to override in the selected stack instances.")
		cloudformation_createStackInstancesCmd.Flags().String("regions", "", "The names of one or more Amazon Web Services Regions where you want to create stack instances using the specified Amazon Web Services accounts.")
		cloudformation_createStackInstancesCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet that you want to create stack instances from.")
		cloudformation_createStackInstancesCmd.MarkFlagRequired("regions")
		cloudformation_createStackInstancesCmd.MarkFlagRequired("stack-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_createStackInstancesCmd)
}
