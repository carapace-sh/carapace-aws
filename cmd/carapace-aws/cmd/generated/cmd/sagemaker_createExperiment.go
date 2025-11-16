package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createExperimentCmd = &cobra.Command{
	Use:   "create-experiment",
	Short: "Creates a SageMaker *experiment*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createExperimentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createExperimentCmd).Standalone()

		sagemaker_createExperimentCmd.Flags().String("description", "", "The description of the experiment.")
		sagemaker_createExperimentCmd.Flags().String("display-name", "", "The name of the experiment as displayed.")
		sagemaker_createExperimentCmd.Flags().String("experiment-name", "", "The name of the experiment.")
		sagemaker_createExperimentCmd.Flags().String("tags", "", "A list of tags to associate with the experiment.")
		sagemaker_createExperimentCmd.MarkFlagRequired("experiment-name")
	})
	sagemakerCmd.AddCommand(sagemaker_createExperimentCmd)
}
