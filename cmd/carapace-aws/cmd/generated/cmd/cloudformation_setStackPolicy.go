package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_setStackPolicyCmd = &cobra.Command{
	Use:   "set-stack-policy",
	Short: "Sets a stack policy for a specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_setStackPolicyCmd).Standalone()

	cloudformation_setStackPolicyCmd.Flags().String("stack-name", "", "The name or unique stack ID that you want to associate a policy with.")
	cloudformation_setStackPolicyCmd.Flags().String("stack-policy-body", "", "Structure that contains the stack policy body.")
	cloudformation_setStackPolicyCmd.Flags().String("stack-policy-url", "", "Location of a file that contains the stack policy.")
	cloudformation_setStackPolicyCmd.MarkFlagRequired("stack-name")
	cloudformationCmd.AddCommand(cloudformation_setStackPolicyCmd)
}
