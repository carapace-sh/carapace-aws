package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_startNotebookInstanceCmd = &cobra.Command{
	Use:   "start-notebook-instance",
	Short: "Launches an ML compute instance with the latest version of the libraries and attaches your ML storage volume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_startNotebookInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_startNotebookInstanceCmd).Standalone()

		sagemaker_startNotebookInstanceCmd.Flags().String("notebook-instance-name", "", "The name of the notebook instance to start.")
		sagemaker_startNotebookInstanceCmd.MarkFlagRequired("notebook-instance-name")
	})
	sagemakerCmd.AddCommand(sagemaker_startNotebookInstanceCmd)
}
