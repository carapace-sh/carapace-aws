package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopTransformJobCmd = &cobra.Command{
	Use:   "stop-transform-job",
	Short: "Stops a batch transform job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopTransformJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_stopTransformJobCmd).Standalone()

		sagemaker_stopTransformJobCmd.Flags().String("transform-job-name", "", "The name of the batch transform job to stop.")
		sagemaker_stopTransformJobCmd.MarkFlagRequired("transform-job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_stopTransformJobCmd)
}
