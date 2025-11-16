package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteOptimizationJobCmd = &cobra.Command{
	Use:   "delete-optimization-job",
	Short: "Deletes an optimization job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteOptimizationJobCmd).Standalone()

	sagemaker_deleteOptimizationJobCmd.Flags().String("optimization-job-name", "", "The name that you assigned to the optimization job.")
	sagemaker_deleteOptimizationJobCmd.MarkFlagRequired("optimization-job-name")
	sagemakerCmd.AddCommand(sagemaker_deleteOptimizationJobCmd)
}
