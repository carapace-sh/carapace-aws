package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listMlflowTrackingServersCmd = &cobra.Command{
	Use:   "list-mlflow-tracking-servers",
	Short: "Lists all MLflow Tracking Servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listMlflowTrackingServersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listMlflowTrackingServersCmd).Standalone()

		sagemaker_listMlflowTrackingServersCmd.Flags().String("created-after", "", "Use the `CreatedAfter` filter to only list tracking servers created after a specific date and time.")
		sagemaker_listMlflowTrackingServersCmd.Flags().String("created-before", "", "Use the `CreatedBefore` filter to only list tracking servers created before a specific date and time.")
		sagemaker_listMlflowTrackingServersCmd.Flags().String("max-results", "", "The maximum number of tracking servers to list.")
		sagemaker_listMlflowTrackingServersCmd.Flags().String("mlflow-version", "", "Filter for tracking servers using the specified MLflow version.")
		sagemaker_listMlflowTrackingServersCmd.Flags().String("next-token", "", "If the previous response was truncated, you will receive this token.")
		sagemaker_listMlflowTrackingServersCmd.Flags().String("sort-by", "", "Filter for trackings servers sorting by name, creation time, or creation status.")
		sagemaker_listMlflowTrackingServersCmd.Flags().String("sort-order", "", "Change the order of the listed tracking servers.")
		sagemaker_listMlflowTrackingServersCmd.Flags().String("tracking-server-status", "", "Filter for tracking servers with a specified creation status.")
	})
	sagemakerCmd.AddCommand(sagemaker_listMlflowTrackingServersCmd)
}
