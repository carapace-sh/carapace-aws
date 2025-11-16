package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_createComponentCmd = &cobra.Command{
	Use:   "create-component",
	Short: "Creates a new component that can be used to build, validate, test, and assess your image.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_createComponentCmd).Standalone()

	imagebuilder_createComponentCmd.Flags().String("change-description", "", "The change description of the component.")
	imagebuilder_createComponentCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
	imagebuilder_createComponentCmd.Flags().String("data", "", "Component `data` contains inline YAML document content for the component.")
	imagebuilder_createComponentCmd.Flags().String("description", "", "Describes the contents of the component.")
	imagebuilder_createComponentCmd.Flags().String("kms-key-id", "", "The Amazon Resource Name (ARN) that uniquely identifies the KMS key used to encrypt this component.")
	imagebuilder_createComponentCmd.Flags().String("name", "", "The name of the component.")
	imagebuilder_createComponentCmd.Flags().String("platform", "", "The operating system platform of the component.")
	imagebuilder_createComponentCmd.Flags().String("semantic-version", "", "The semantic version of the component.")
	imagebuilder_createComponentCmd.Flags().String("supported-os-versions", "", "The operating system (OS) version supported by the component.")
	imagebuilder_createComponentCmd.Flags().String("tags", "", "The tags that apply to the component.")
	imagebuilder_createComponentCmd.Flags().String("uri", "", "The `uri` of a YAML component document file.")
	imagebuilder_createComponentCmd.MarkFlagRequired("client-token")
	imagebuilder_createComponentCmd.MarkFlagRequired("name")
	imagebuilder_createComponentCmd.MarkFlagRequired("platform")
	imagebuilder_createComponentCmd.MarkFlagRequired("semantic-version")
	imagebuilderCmd.AddCommand(imagebuilder_createComponentCmd)
}
