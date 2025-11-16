package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listStackInstancesCmd = &cobra.Command{
	Use:   "list-stack-instances",
	Short: "Returns summary information about stack instances that are associated with the specified StackSet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listStackInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listStackInstancesCmd).Standalone()

		cloudformation_listStackInstancesCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
		cloudformation_listStackInstancesCmd.Flags().String("filters", "", "The filter to apply to stack instances")
		cloudformation_listStackInstancesCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
		cloudformation_listStackInstancesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listStackInstancesCmd.Flags().String("stack-instance-account", "", "The name of the Amazon Web Services account that you want to list stack instances for.")
		cloudformation_listStackInstancesCmd.Flags().String("stack-instance-region", "", "The name of the Region where you want to list stack instances.")
		cloudformation_listStackInstancesCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet that you want to list stack instances for.")
		cloudformation_listStackInstancesCmd.MarkFlagRequired("stack-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_listStackInstancesCmd)
}
