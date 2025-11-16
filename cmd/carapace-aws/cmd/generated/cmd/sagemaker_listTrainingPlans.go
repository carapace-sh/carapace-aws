package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listTrainingPlansCmd = &cobra.Command{
	Use:   "list-training-plans",
	Short: "Retrieves a list of training plans for the current account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listTrainingPlansCmd).Standalone()

	sagemaker_listTrainingPlansCmd.Flags().String("filters", "", "Additional filters to apply to the list of training plans.")
	sagemaker_listTrainingPlansCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	sagemaker_listTrainingPlansCmd.Flags().String("next-token", "", "A token to continue pagination if more results are available.")
	sagemaker_listTrainingPlansCmd.Flags().String("sort-by", "", "The training plan field to sort the results by (e.g., StartTime, Status).")
	sagemaker_listTrainingPlansCmd.Flags().String("sort-order", "", "The order to sort the results (Ascending or Descending).")
	sagemaker_listTrainingPlansCmd.Flags().String("start-time-after", "", "Filter to list only training plans with an actual start time after this date.")
	sagemaker_listTrainingPlansCmd.Flags().String("start-time-before", "", "Filter to list only training plans with an actual start time before this date.")
	sagemakerCmd.AddCommand(sagemaker_listTrainingPlansCmd)
}
