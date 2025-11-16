package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getImageRecipePolicyCmd = &cobra.Command{
	Use:   "get-image-recipe-policy",
	Short: "Gets an image recipe policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getImageRecipePolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_getImageRecipePolicyCmd).Standalone()

		imagebuilder_getImageRecipePolicyCmd.Flags().String("image-recipe-arn", "", "The Amazon Resource Name (ARN) of the image recipe whose policy you want to retrieve.")
		imagebuilder_getImageRecipePolicyCmd.MarkFlagRequired("image-recipe-arn")
	})
	imagebuilderCmd.AddCommand(imagebuilder_getImageRecipePolicyCmd)
}
