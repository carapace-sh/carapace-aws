package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createMlflowTrackingServerCmd = &cobra.Command{
	Use:   "create-mlflow-tracking-server",
	Short: "Creates an MLflow Tracking Server using a general purpose Amazon S3 bucket as the artifact store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createMlflowTrackingServerCmd).Standalone()

	sagemaker_createMlflowTrackingServerCmd.Flags().String("artifact-store-uri", "", "The S3 URI for a general purpose bucket to use as the MLflow Tracking Server artifact store.")
	sagemaker_createMlflowTrackingServerCmd.Flags().Bool("automatic-model-registration", false, "Whether to enable or disable automatic registration of new MLflow models to the SageMaker Model Registry.")
	sagemaker_createMlflowTrackingServerCmd.Flags().String("mlflow-version", "", "The version of MLflow that the tracking server uses.")
	sagemaker_createMlflowTrackingServerCmd.Flags().Bool("no-automatic-model-registration", false, "Whether to enable or disable automatic registration of new MLflow models to the SageMaker Model Registry.")
	sagemaker_createMlflowTrackingServerCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) for an IAM role in your account that the MLflow Tracking Server uses to access the artifact store in Amazon S3.")
	sagemaker_createMlflowTrackingServerCmd.Flags().String("tags", "", "Tags consisting of key-value pairs used to manage metadata for the tracking server.")
	sagemaker_createMlflowTrackingServerCmd.Flags().String("tracking-server-name", "", "A unique string identifying the tracking server name.")
	sagemaker_createMlflowTrackingServerCmd.Flags().String("tracking-server-size", "", "The size of the tracking server you want to create.")
	sagemaker_createMlflowTrackingServerCmd.Flags().String("weekly-maintenance-window-start", "", "The day and time of the week in Coordinated Universal Time (UTC) 24-hour standard time that weekly maintenance updates are scheduled.")
	sagemaker_createMlflowTrackingServerCmd.MarkFlagRequired("artifact-store-uri")
	sagemaker_createMlflowTrackingServerCmd.Flag("no-automatic-model-registration").Hidden = true
	sagemaker_createMlflowTrackingServerCmd.MarkFlagRequired("role-arn")
	sagemaker_createMlflowTrackingServerCmd.MarkFlagRequired("tracking-server-name")
	sagemakerCmd.AddCommand(sagemaker_createMlflowTrackingServerCmd)
}
