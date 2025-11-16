package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateArtifactCmd = &cobra.Command{
	Use:   "update-artifact",
	Short: "Updates an artifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateArtifactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateArtifactCmd).Standalone()

		sagemaker_updateArtifactCmd.Flags().String("artifact-arn", "", "The Amazon Resource Name (ARN) of the artifact to update.")
		sagemaker_updateArtifactCmd.Flags().String("artifact-name", "", "The new name for the artifact.")
		sagemaker_updateArtifactCmd.Flags().String("properties", "", "The new list of properties.")
		sagemaker_updateArtifactCmd.Flags().String("properties-to-remove", "", "A list of properties to remove.")
		sagemaker_updateArtifactCmd.MarkFlagRequired("artifact-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_updateArtifactCmd)
}
