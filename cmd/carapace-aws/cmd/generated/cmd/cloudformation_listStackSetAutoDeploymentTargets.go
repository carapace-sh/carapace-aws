package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_listStackSetAutoDeploymentTargetsCmd = &cobra.Command{
	Use:   "list-stack-set-auto-deployment-targets",
	Short: "Returns summary information about deployment targets for a StackSet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_listStackSetAutoDeploymentTargetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_listStackSetAutoDeploymentTargetsCmd).Standalone()

		cloudformation_listStackSetAutoDeploymentTargetsCmd.Flags().String("call-as", "", "Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
		cloudformation_listStackSetAutoDeploymentTargetsCmd.Flags().String("max-results", "", "The maximum number of results to be returned with a single call.")
		cloudformation_listStackSetAutoDeploymentTargetsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
		cloudformation_listStackSetAutoDeploymentTargetsCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet that you want to get automatic deployment targets for.")
		cloudformation_listStackSetAutoDeploymentTargetsCmd.MarkFlagRequired("stack-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_listStackSetAutoDeploymentTargetsCmd)
}
