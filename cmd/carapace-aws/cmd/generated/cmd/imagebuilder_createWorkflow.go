package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var imagebuilder_createWorkflowCmd = &cobra.Command{
	Use:   "create-workflow",
	Short: "Create a new workflow or a new version of an existing workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagebuilder_createWorkflowCmd).Standalone()

	imagebuilder_createWorkflowCmd.Flags().String("change-description", "", "Describes what change has been made in this version of the workflow, or what makes this version different from other versions of the workflow.")
	imagebuilder_createWorkflowCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure idempotency of the request.")
	imagebuilder_createWorkflowCmd.Flags().String("data", "", "Contains the UTF-8 encoded YAML document content for the workflow.")
	imagebuilder_createWorkflowCmd.Flags().String("description", "", "Describes the workflow.")
	imagebuilder_createWorkflowCmd.Flags().String("kms-key-id", "", "The Amazon Resource Name (ARN) that uniquely identifies the KMS key used to encrypt this workflow resource.")
	imagebuilder_createWorkflowCmd.Flags().String("name", "", "The name of the workflow to create.")
	imagebuilder_createWorkflowCmd.Flags().String("semantic-version", "", "The semantic version of this workflow resource.")
	imagebuilder_createWorkflowCmd.Flags().String("tags", "", "Tags that apply to the workflow resource.")
	imagebuilder_createWorkflowCmd.Flags().String("type", "", "The phase in the image build process for which the workflow resource is responsible.")
	imagebuilder_createWorkflowCmd.Flags().String("uri", "", "The `uri` of a YAML component document file.")
	imagebuilder_createWorkflowCmd.MarkFlagRequired("client-token")
	imagebuilder_createWorkflowCmd.MarkFlagRequired("name")
	imagebuilder_createWorkflowCmd.MarkFlagRequired("semantic-version")
	imagebuilder_createWorkflowCmd.MarkFlagRequired("type")
	imagebuilderCmd.AddCommand(imagebuilder_createWorkflowCmd)
}
