package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_getImagePolicyCmd = &cobra.Command{
	Use:   "get-image-policy",
	Short: "Gets an image policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_getImagePolicyCmd).Standalone()

	imagebuilder_getImagePolicyCmd.Flags().String("image-arn", "", "The Amazon Resource Name (ARN) of the image whose policy you want to retrieve.")
	imagebuilder_getImagePolicyCmd.MarkFlagRequired("image-arn")
	imagebuilderCmd.AddCommand(imagebuilder_getImagePolicyCmd)
}
