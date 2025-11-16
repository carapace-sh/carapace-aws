package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeExperimentCmd = &cobra.Command{
	Use:   "describe-experiment",
	Short: "Provides a list of an experiment's properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeExperimentCmd).Standalone()

	sagemaker_describeExperimentCmd.Flags().String("experiment-name", "", "The name of the experiment to describe.")
	sagemaker_describeExperimentCmd.MarkFlagRequired("experiment-name")
	sagemakerCmd.AddCommand(sagemaker_describeExperimentCmd)
}
