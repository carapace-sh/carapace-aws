package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateExperimentCmd = &cobra.Command{
	Use:   "update-experiment",
	Short: "Adds, updates, or removes the description of an experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateExperimentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_updateExperimentCmd).Standalone()

		sagemaker_updateExperimentCmd.Flags().String("description", "", "The description of the experiment.")
		sagemaker_updateExperimentCmd.Flags().String("display-name", "", "The name of the experiment as displayed.")
		sagemaker_updateExperimentCmd.Flags().String("experiment-name", "", "The name of the experiment to update.")
		sagemaker_updateExperimentCmd.MarkFlagRequired("experiment-name")
	})
	sagemakerCmd.AddCommand(sagemaker_updateExperimentCmd)
}
