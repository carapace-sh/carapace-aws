package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteArtifactCmd = &cobra.Command{
	Use:   "delete-artifact",
	Short: "Deletes an artifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteArtifactCmd).Standalone()

	sagemaker_deleteArtifactCmd.Flags().String("artifact-arn", "", "The Amazon Resource Name (ARN) of the artifact to delete.")
	sagemaker_deleteArtifactCmd.Flags().String("source", "", "The URI of the source.")
	sagemakerCmd.AddCommand(sagemaker_deleteArtifactCmd)
}
