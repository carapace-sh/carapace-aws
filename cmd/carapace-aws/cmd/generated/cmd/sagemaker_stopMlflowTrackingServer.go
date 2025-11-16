package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_stopMlflowTrackingServerCmd = &cobra.Command{
	Use:   "stop-mlflow-tracking-server",
	Short: "Programmatically stop an MLflow Tracking Server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_stopMlflowTrackingServerCmd).Standalone()

	sagemaker_stopMlflowTrackingServerCmd.Flags().String("tracking-server-name", "", "The name of the tracking server to stop.")
	sagemaker_stopMlflowTrackingServerCmd.MarkFlagRequired("tracking-server-name")
	sagemakerCmd.AddCommand(sagemaker_stopMlflowTrackingServerCmd)
}
