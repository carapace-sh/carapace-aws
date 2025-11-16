package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeArtifactCmd = &cobra.Command{
	Use:   "describe-artifact",
	Short: "Describes an artifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeArtifactCmd).Standalone()

	sagemaker_describeArtifactCmd.Flags().String("artifact-arn", "", "The Amazon Resource Name (ARN) of the artifact to describe.")
	sagemaker_describeArtifactCmd.MarkFlagRequired("artifact-arn")
	sagemakerCmd.AddCommand(sagemaker_describeArtifactCmd)
}
