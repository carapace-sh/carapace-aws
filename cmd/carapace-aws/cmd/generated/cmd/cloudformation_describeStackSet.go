package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeStackSetCmd = &cobra.Command{
	Use:   "describe-stack-set",
	Short: "Returns the description of the specified StackSet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeStackSetCmd).Standalone()

	cloudformation_describeStackSetCmd.Flags().String("call-as", "", "\\[Service-managed permissions] Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account.")
	cloudformation_describeStackSetCmd.Flags().String("stack-set-name", "", "The name or unique ID of the StackSet whose description you want.")
	cloudformation_describeStackSetCmd.MarkFlagRequired("stack-set-name")
	cloudformationCmd.AddCommand(cloudformation_describeStackSetCmd)
}
