package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_putImageRecipePolicyCmd = &cobra.Command{
	Use:   "put-image-recipe-policy",
	Short: "Applies a policy to an image recipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_putImageRecipePolicyCmd).Standalone()

	imagebuilder_putImageRecipePolicyCmd.Flags().String("image-recipe-arn", "", "The Amazon Resource Name (ARN) of the image recipe that this policy should be applied to.")
	imagebuilder_putImageRecipePolicyCmd.Flags().String("policy", "", "The policy to apply.")
	imagebuilder_putImageRecipePolicyCmd.MarkFlagRequired("image-recipe-arn")
	imagebuilder_putImageRecipePolicyCmd.MarkFlagRequired("policy")
	imagebuilderCmd.AddCommand(imagebuilder_putImageRecipePolicyCmd)
}
