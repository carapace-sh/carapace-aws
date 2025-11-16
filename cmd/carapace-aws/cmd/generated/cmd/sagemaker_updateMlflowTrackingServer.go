package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_updateMlflowTrackingServerCmd = &cobra.Command{
	Use:   "update-mlflow-tracking-server",
	Short: "Updates properties of an existing MLflow Tracking Server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_updateMlflowTrackingServerCmd).Standalone()

	sagemaker_updateMlflowTrackingServerCmd.Flags().String("artifact-store-uri", "", "The new S3 URI for the general purpose bucket to use as the artifact store for the MLflow Tracking Server.")
	sagemaker_updateMlflowTrackingServerCmd.Flags().Bool("automatic-model-registration", false, "Whether to enable or disable automatic registration of new MLflow models to the SageMaker Model Registry.")
	sagemaker_updateMlflowTrackingServerCmd.Flags().Bool("no-automatic-model-registration", false, "Whether to enable or disable automatic registration of new MLflow models to the SageMaker Model Registry.")
	sagemaker_updateMlflowTrackingServerCmd.Flags().String("tracking-server-name", "", "The name of the MLflow Tracking Server to update.")
	sagemaker_updateMlflowTrackingServerCmd.Flags().String("tracking-server-size", "", "The new size for the MLflow Tracking Server.")
	sagemaker_updateMlflowTrackingServerCmd.Flags().String("weekly-maintenance-window-start", "", "The new weekly maintenance window start day and time to update.")
	sagemaker_updateMlflowTrackingServerCmd.Flag("no-automatic-model-registration").Hidden = true
	sagemaker_updateMlflowTrackingServerCmd.MarkFlagRequired("tracking-server-name")
	sagemakerCmd.AddCommand(sagemaker_updateMlflowTrackingServerCmd)
}
