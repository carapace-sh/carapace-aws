package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopOptimizationJobCmd = &cobra.Command{
	Use:   "stop-optimization-job",
	Short: "Ends a running inference optimization job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopOptimizationJobCmd).Standalone()

	sagemaker_stopOptimizationJobCmd.Flags().String("optimization-job-name", "", "The name that you assigned to the optimization job.")
	sagemaker_stopOptimizationJobCmd.MarkFlagRequired("optimization-job-name")
	sagemakerCmd.AddCommand(sagemaker_stopOptimizationJobCmd)
}
