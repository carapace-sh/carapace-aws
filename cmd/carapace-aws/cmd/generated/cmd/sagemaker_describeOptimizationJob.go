package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeOptimizationJobCmd = &cobra.Command{
	Use:   "describe-optimization-job",
	Short: "Provides the properties of the specified optimization job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeOptimizationJobCmd).Standalone()

	sagemaker_describeOptimizationJobCmd.Flags().String("optimization-job-name", "", "The name that you assigned to the optimization job.")
	sagemaker_describeOptimizationJobCmd.MarkFlagRequired("optimization-job-name")
	sagemakerCmd.AddCommand(sagemaker_describeOptimizationJobCmd)
}
