package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeStackInstanceCmd = &cobra.Command{
	Use:   "describe-stack-instance",
	Short: "Returns the stack instance that's associated with the specified StackSet, Amazon Web Services account, and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeStackInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_describeStackInstanceCmd).Standalone()

		cloudformation_describeStackInstanceCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
		cloudformation_describeStackInstanceCmd.Flags().String("stack-instance-account", "", "The ID of an Amazon Web Services account that's associated with this stack instance.")
		cloudformation_describeStackInstanceCmd.Flags().String("stack-instance-region", "", "The name of a Region that's associated with this stack instance.")
		cloudformation_describeStackInstanceCmd.Flags().String("stack-set-name", "", "The name or the unique stack ID of the StackSet that you want to get stack instance information for.")
		cloudformation_describeStackInstanceCmd.MarkFlagRequired("stack-instance-account")
		cloudformation_describeStackInstanceCmd.MarkFlagRequired("stack-instance-region")
		cloudformation_describeStackInstanceCmd.MarkFlagRequired("stack-set-name")
	})
	cloudformationCmd.AddCommand(cloudformation_describeStackInstanceCmd)
}
