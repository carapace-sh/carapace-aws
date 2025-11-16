package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeTransformJobCmd = &cobra.Command{
	Use:   "describe-transform-job",
	Short: "Returns information about a transform job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeTransformJobCmd).Standalone()

	sagemaker_describeTransformJobCmd.Flags().String("transform-job-name", "", "The name of the transform job that you want to view details of.")
	sagemaker_describeTransformJobCmd.MarkFlagRequired("transform-job-name")
	sagemakerCmd.AddCommand(sagemaker_describeTransformJobCmd)
}
