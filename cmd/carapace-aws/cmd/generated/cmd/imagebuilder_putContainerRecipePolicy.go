package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_putContainerRecipePolicyCmd = &cobra.Command{
	Use:   "put-container-recipe-policy",
	Short: "Applies a policy to a container image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_putContainerRecipePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_putContainerRecipePolicyCmd).Standalone()

		imagebuilder_putContainerRecipePolicyCmd.Flags().String("container-recipe-arn", "", "The Amazon Resource Name (ARN) of the container recipe that this policy should be applied to.")
		imagebuilder_putContainerRecipePolicyCmd.Flags().String("policy", "", "The policy to apply to the container recipe.")
		imagebuilder_putContainerRecipePolicyCmd.MarkFlagRequired("container-recipe-arn")
		imagebuilder_putContainerRecipePolicyCmd.MarkFlagRequired("policy")
	})
	imagebuilderCmd.AddCommand(imagebuilder_putContainerRecipePolicyCmd)
}
