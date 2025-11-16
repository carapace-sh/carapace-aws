package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_searchTrainingPlanOfferingsCmd = &cobra.Command{
	Use:   "search-training-plan-offerings",
	Short: "Searches for available training plan offerings based on specified criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_searchTrainingPlanOfferingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_searchTrainingPlanOfferingsCmd).Standalone()

		sagemaker_searchTrainingPlanOfferingsCmd.Flags().String("duration-hours", "", "The desired duration in hours for the training plan offerings.")
		sagemaker_searchTrainingPlanOfferingsCmd.Flags().String("end-time-before", "", "A filter to search for reserved capacity offerings with an end time before a specified date.")
		sagemaker_searchTrainingPlanOfferingsCmd.Flags().String("instance-count", "", "The number of instances you want to reserve in the training plan offerings.")
		sagemaker_searchTrainingPlanOfferingsCmd.Flags().String("instance-type", "", "The type of instance you want to search for in the available training plan offerings.")
		sagemaker_searchTrainingPlanOfferingsCmd.Flags().String("start-time-after", "", "A filter to search for training plan offerings with a start time after a specified date.")
		sagemaker_searchTrainingPlanOfferingsCmd.Flags().String("target-resources", "", "The target resources (e.g., SageMaker Training Jobs, SageMaker HyperPod) to search for in the offerings.")
		sagemaker_searchTrainingPlanOfferingsCmd.Flags().String("ultra-server-count", "", "The number of UltraServers to search for.")
		sagemaker_searchTrainingPlanOfferingsCmd.Flags().String("ultra-server-type", "", "The type of UltraServer to search for, such as ml.u-p6e-gb200x72.")
		sagemaker_searchTrainingPlanOfferingsCmd.MarkFlagRequired("duration-hours")
		sagemaker_searchTrainingPlanOfferingsCmd.MarkFlagRequired("target-resources")
	})
	sagemakerCmd.AddCommand(sagemaker_searchTrainingPlanOfferingsCmd)
}
