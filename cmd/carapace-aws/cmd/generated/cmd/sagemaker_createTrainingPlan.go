package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createTrainingPlanCmd = &cobra.Command{
	Use:   "create-training-plan",
	Short: "Creates a new training plan in SageMaker to reserve compute capacity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createTrainingPlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createTrainingPlanCmd).Standalone()

		sagemaker_createTrainingPlanCmd.Flags().String("spare-instance-count-per-ultra-server", "", "Number of spare instances to reserve per UltraServer for enhanced resiliency.")
		sagemaker_createTrainingPlanCmd.Flags().String("tags", "", "An array of key-value pairs to apply to this training plan.")
		sagemaker_createTrainingPlanCmd.Flags().String("training-plan-name", "", "The name of the training plan to create.")
		sagemaker_createTrainingPlanCmd.Flags().String("training-plan-offering-id", "", "The unique identifier of the training plan offering to use for creating this plan.")
		sagemaker_createTrainingPlanCmd.MarkFlagRequired("training-plan-name")
		sagemaker_createTrainingPlanCmd.MarkFlagRequired("training-plan-offering-id")
	})
	sagemakerCmd.AddCommand(sagemaker_createTrainingPlanCmd)
}
