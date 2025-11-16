package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteExperimentCmd = &cobra.Command{
	Use:   "delete-experiment",
	Short: "Deletes an SageMaker experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteExperimentCmd).Standalone()

	sagemaker_deleteExperimentCmd.Flags().String("experiment-name", "", "The name of the experiment to delete.")
	sagemaker_deleteExperimentCmd.MarkFlagRequired("experiment-name")
	sagemakerCmd.AddCommand(sagemaker_deleteExperimentCmd)
}
