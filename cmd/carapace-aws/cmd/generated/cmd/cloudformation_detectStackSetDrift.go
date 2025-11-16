package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_detectStackSetDriftCmd = &cobra.Command{
	Use:   "detect-stack-set-drift",
	Short: "Detect drift on a StackSet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_detectStackSetDriftCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_detectStackSetDriftCmd).Standalone()

		cloudformation_detectStackSetDriftCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
		cloudformation_detectStackSetDriftCmd.Flags().String("operation-id", "", "*The ID of the StackSet operation.*")
		cloudformation_detectStackSetDriftCmd.Flags().String("operation-preferences", "", "The user-specified preferences for how CloudFormation performs a StackSet operation.")
		cloudformation_detectStackSetDriftCmd.Flags().String("stack-set-name", "", "The name of the StackSet on which to perform the drift detection operation.")
		cloudformation_detectStackSetDriftCmd.MarkFlagRequired("stack-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_detectStackSetDriftCmd)
}
