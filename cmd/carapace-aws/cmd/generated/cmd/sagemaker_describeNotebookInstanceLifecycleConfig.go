package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeNotebookInstanceLifecycleConfigCmd = &cobra.Command{
	Use:   "describe-notebook-instance-lifecycle-config",
	Short: "Returns a description of a notebook instance lifecycle configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeNotebookInstanceLifecycleConfigCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeNotebookInstanceLifecycleConfigCmd).Standalone()

		sagemaker_describeNotebookInstanceLifecycleConfigCmd.Flags().String("notebook-instance-lifecycle-config-name", "", "The name of the lifecycle configuration to describe.")
		sagemaker_describeNotebookInstanceLifecycleConfigCmd.MarkFlagRequired("notebook-instance-lifecycle-config-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeNotebookInstanceLifecycleConfigCmd)
}
