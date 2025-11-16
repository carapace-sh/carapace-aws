package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeNotebookInstanceCmd = &cobra.Command{
	Use:   "describe-notebook-instance",
	Short: "Returns information about a notebook instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeNotebookInstanceCmd).Standalone()

	sagemaker_describeNotebookInstanceCmd.Flags().String("notebook-instance-name", "", "The name of the notebook instance that you want information about.")
	sagemaker_describeNotebookInstanceCmd.MarkFlagRequired("notebook-instance-name")
	sagemakerCmd.AddCommand(sagemaker_describeNotebookInstanceCmd)
}
