package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_cancelImageCreationCmd = &cobra.Command{
	Use:   "cancel-image-creation",
	Short: "CancelImageCreation cancels the creation of Image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_cancelImageCreationCmd).Standalone()

	imagebuilder_cancelImageCreationCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
	imagebuilder_cancelImageCreationCmd.Flags().String("image-build-version-arn", "", "The Amazon Resource Name (ARN) of the image that you want to cancel creation for.")
	imagebuilder_cancelImageCreationCmd.MarkFlagRequired("client-token")
	imagebuilder_cancelImageCreationCmd.MarkFlagRequired("image-build-version-arn")
	imagebuilderCmd.AddCommand(imagebuilder_cancelImageCreationCmd)
}
