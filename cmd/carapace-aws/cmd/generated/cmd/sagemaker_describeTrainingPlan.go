package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeTrainingPlanCmd = &cobra.Command{
	Use:   "describe-training-plan",
	Short: "Retrieves detailed information about a specific training plan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeTrainingPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_describeTrainingPlanCmd).Standalone()

		sagemaker_describeTrainingPlanCmd.Flags().String("training-plan-name", "", "The name of the training plan to describe.")
		sagemaker_describeTrainingPlanCmd.MarkFlagRequired("training-plan-name")
	})
	sagemakerCmd.AddCommand(sagemaker_describeTrainingPlanCmd)
}
