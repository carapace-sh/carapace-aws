package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_deleteMlflowTrackingServerCmd = &cobra.Command{
	Use:   "delete-mlflow-tracking-server",
	Short: "Deletes an MLflow Tracking Server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_deleteMlflowTrackingServerCmd).Standalone()

	sagemaker_deleteMlflowTrackingServerCmd.Flags().String("tracking-server-name", "", "The name of the the tracking server to delete.")
	sagemaker_deleteMlflowTrackingServerCmd.MarkFlagRequired("tracking-server-name")
	sagemakerCmd.AddCommand(sagemaker_deleteMlflowTrackingServerCmd)
}
