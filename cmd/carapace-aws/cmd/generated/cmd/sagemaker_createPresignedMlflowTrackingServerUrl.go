package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createPresignedMlflowTrackingServerUrlCmd = &cobra.Command{
	Use:   "create-presigned-mlflow-tracking-server-url",
	Short: "Returns a presigned URL that you can use to connect to the MLflow UI attached to your tracking server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createPresignedMlflowTrackingServerUrlCmd).Standalone()

	sagemaker_createPresignedMlflowTrackingServerUrlCmd.Flags().String("expires-in-seconds", "", "The duration in seconds that your presigned URL is valid.")
	sagemaker_createPresignedMlflowTrackingServerUrlCmd.Flags().String("session-expiration-duration-in-seconds", "", "The duration in seconds that your MLflow UI session is valid.")
	sagemaker_createPresignedMlflowTrackingServerUrlCmd.Flags().String("tracking-server-name", "", "The name of the tracking server to connect to your MLflow UI.")
	sagemaker_createPresignedMlflowTrackingServerUrlCmd.MarkFlagRequired("tracking-server-name")
	sagemakerCmd.AddCommand(sagemaker_createPresignedMlflowTrackingServerUrlCmd)
}
