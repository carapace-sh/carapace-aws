package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_importComponentCmd = &cobra.Command{
	Use:   "import-component",
	Short: "Imports a component and transforms its data into a component document.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_importComponentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(imagebuilder_importComponentCmd).Standalone()

		imagebuilder_importComponentCmd.Flags().String("change-description", "", "The change description of the component.")
		imagebuilder_importComponentCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
		imagebuilder_importComponentCmd.Flags().String("data", "", "The data of the component.")
		imagebuilder_importComponentCmd.Flags().String("description", "", "The description of the component.")
		imagebuilder_importComponentCmd.Flags().String("format", "", "The format of the resource that you want to import as a component.")
		imagebuilder_importComponentCmd.Flags().String("kms-key-id", "", "The Amazon Resource Name (ARN) that uniquely identifies the KMS key used to encrypt this component.")
		imagebuilder_importComponentCmd.Flags().String("name", "", "The name of the component.")
		imagebuilder_importComponentCmd.Flags().String("platform", "", "The platform of the component.")
		imagebuilder_importComponentCmd.Flags().String("semantic-version", "", "The semantic version of the component.")
		imagebuilder_importComponentCmd.Flags().String("tags", "", "The tags of the component.")
		imagebuilder_importComponentCmd.Flags().String("type", "", "The type of the component denotes whether the component is used to build the image, or only to test it.")
		imagebuilder_importComponentCmd.Flags().String("uri", "", "The uri of the component.")
		imagebuilder_importComponentCmd.MarkFlagRequired("client-token")
		imagebuilder_importComponentCmd.MarkFlagRequired("format")
		imagebuilder_importComponentCmd.MarkFlagRequired("name")
		imagebuilder_importComponentCmd.MarkFlagRequired("platform")
		imagebuilder_importComponentCmd.MarkFlagRequired("semantic-version")
		imagebuilder_importComponentCmd.MarkFlagRequired("type")
	})
	imagebuilderCmd.AddCommand(imagebuilder_importComponentCmd)
}
