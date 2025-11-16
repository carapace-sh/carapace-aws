package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_startMlflowTrackingServerCmd = &cobra.Command{
	Use:   "start-mlflow-tracking-server",
	Short: "Programmatically start an MLflow Tracking Server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_startMlflowTrackingServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_startMlflowTrackingServerCmd).Standalone()

		sagemaker_startMlflowTrackingServerCmd.Flags().String("tracking-server-name", "", "The name of the tracking server to start.")
		sagemaker_startMlflowTrackingServerCmd.MarkFlagRequired("tracking-server-name")
	})
	sagemakerCmd.AddCommand(sagemaker_startMlflowTrackingServerCmd)
}
