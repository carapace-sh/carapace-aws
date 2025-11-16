package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createArtifactCmd = &cobra.Command{
	Use:   "create-artifact",
	Short: "Creates an *artifact*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createArtifactCmd).Standalone()

	sagemaker_createArtifactCmd.Flags().String("artifact-name", "", "The name of the artifact.")
	sagemaker_createArtifactCmd.Flags().String("artifact-type", "", "The artifact type.")
	sagemaker_createArtifactCmd.Flags().String("metadata-properties", "", "")
	sagemaker_createArtifactCmd.Flags().String("properties", "", "A list of properties to add to the artifact.")
	sagemaker_createArtifactCmd.Flags().String("source", "", "The ID, ID type, and URI of the source.")
	sagemaker_createArtifactCmd.Flags().String("tags", "", "A list of tags to apply to the artifact.")
	sagemaker_createArtifactCmd.MarkFlagRequired("artifact-type")
	sagemaker_createArtifactCmd.MarkFlagRequired("source")
	sagemakerCmd.AddCommand(sagemaker_createArtifactCmd)
}
