package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_putComponentPolicyCmd = &cobra.Command{
	Use:   "put-component-policy",
	Short: "Applies a policy to a component.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_putComponentPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_putComponentPolicyCmd).Standalone()

		imagebuilder_putComponentPolicyCmd.Flags().String("component-arn", "", "The Amazon Resource Name (ARN) of the component that this policy should be applied to.")
		imagebuilder_putComponentPolicyCmd.Flags().String("policy", "", "The policy to apply.")
		imagebuilder_putComponentPolicyCmd.MarkFlagRequired("component-arn")
		imagebuilder_putComponentPolicyCmd.MarkFlagRequired("policy")
	})
	imagebuilderCmd.AddCommand(imagebuilder_putComponentPolicyCmd)
}
