package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeCompilationJobCmd = &cobra.Command{
	Use:   "describe-compilation-job",
	Short: "Returns information about a model compilation job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeCompilationJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeCompilationJobCmd).Standalone()

		sagemaker_describeCompilationJobCmd.Flags().String("compilation-job-name", "", "The name of the model compilation job that you want information about.")
		sagemaker_describeCompilationJobCmd.MarkFlagRequired("compilation-job-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeCompilationJobCmd)
}
