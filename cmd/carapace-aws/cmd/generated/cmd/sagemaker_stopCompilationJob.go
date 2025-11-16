package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopCompilationJobCmd = &cobra.Command{
	Use:   "stop-compilation-job",
	Short: "Stops a model compilation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopCompilationJobCmd).Standalone()

	sagemaker_stopCompilationJobCmd.Flags().String("compilation-job-name", "", "The name of the model compilation job to stop.")
	sagemaker_stopCompilationJobCmd.MarkFlagRequired("compilation-job-name")
	sagemakerCmd.AddCommand(sagemaker_stopCompilationJobCmd)
}
