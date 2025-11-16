package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteCompilationJobCmd = &cobra.Command{
	Use:   "delete-compilation-job",
	Short: "Deletes the specified compilation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteCompilationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_deleteCompilationJobCmd).Standalone()

		sagemaker_deleteCompilationJobCmd.Flags().String("compilation-job-name", "", "The name of the compilation job to delete.")
		sagemaker_deleteCompilationJobCmd.MarkFlagRequired("compilation-job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_deleteCompilationJobCmd)
}
