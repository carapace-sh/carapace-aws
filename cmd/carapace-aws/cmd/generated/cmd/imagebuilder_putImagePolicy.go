package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_putImagePolicyCmd = &cobra.Command{
	Use:   "put-image-policy",
	Short: "Applies a policy to an image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_putImagePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_putImagePolicyCmd).Standalone()

		imagebuilder_putImagePolicyCmd.Flags().String("image-arn", "", "The Amazon Resource Name (ARN) of the image that this policy should be applied to.")
		imagebuilder_putImagePolicyCmd.Flags().String("policy", "", "The policy to apply.")
		imagebuilder_putImagePolicyCmd.MarkFlagRequired("image-arn")
		imagebuilder_putImagePolicyCmd.MarkFlagRequired("policy")
	})
	imagebuilderCmd.AddCommand(imagebuilder_putImagePolicyCmd)
}
