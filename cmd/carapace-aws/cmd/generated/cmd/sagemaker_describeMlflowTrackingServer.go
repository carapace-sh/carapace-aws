package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_describeMlflowTrackingServerCmd = &cobra.Command{
	Use:   "describe-mlflow-tracking-server",
	Short: "Returns information about an MLflow Tracking Server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_describeMlflowTrackingServerCmd).Standalone()

	sagemaker_describeMlflowTrackingServerCmd.Flags().String("tracking-server-name", "", "The name of the MLflow Tracking Server to describe.")
	sagemaker_describeMlflowTrackingServerCmd.MarkFlagRequired("tracking-server-name")
	sagemakerCmd.AddCommand(sagemaker_describeMlflowTrackingServerCmd)
}
