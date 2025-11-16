package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopNotebookInstanceCmd = &cobra.Command{
	Use:   "stop-notebook-instance",
	Short: "Terminates the ML compute instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopNotebookInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_stopNotebookInstanceCmd).Standalone()

		sagemaker_stopNotebookInstanceCmd.Flags().String("notebook-instance-name", "", "The name of the notebook instance to terminate.")
		sagemaker_stopNotebookInstanceCmd.MarkFlagRequired("notebook-instance-name")
	})
	sagemakerCmd.AddCommand(sagemaker_stopNotebookInstanceCmd)
}
