package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_getStackPolicyCmd = &cobra.Command{
	Use:   "get-stack-policy",
	Short: "Returns the stack policy for a specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_getStackPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_getStackPolicyCmd).Standalone()

		cloudformation_getStackPolicyCmd.Flags().String("stack-name", "", "The name or unique stack ID that's associated with the stack whose policy you want to get.")
		cloudformation_getStackPolicyCmd.MarkFlagRequired("stack-name")
	})
	cloudformationCmd.AddCommand(cloudformation_getStackPolicyCmd)
}
