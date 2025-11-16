package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_putAnomalyDetectorCmd = &cobra.Command{
	Use:   "put-anomaly-detector",
	Short: "Creates an anomaly detection model for a CloudWatch metric.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_putAnomalyDetectorCmd).Standalone()

	cloudwatch_putAnomalyDetectorCmd.Flags().String("configuration", "", "The configuration specifies details about how the anomaly detection model is to be trained, including time ranges to exclude when training and updating the model.")
	cloudwatch_putAnomalyDetectorCmd.Flags().String("dimensions", "", "The metric dimensions to create the anomaly detection model for.")
	cloudwatch_putAnomalyDetectorCmd.Flags().String("metric-characteristics", "", "Use this object to include parameters to provide information about your metric to CloudWatch to help it build more accurate anomaly detection models.")
	cloudwatch_putAnomalyDetectorCmd.Flags().String("metric-math-anomaly-detector", "", "The metric math anomaly detector to be created.")
	cloudwatch_putAnomalyDetectorCmd.Flags().String("metric-name", "", "The name of the metric to create the anomaly detection model for.")
	cloudwatch_putAnomalyDetectorCmd.Flags().String("namespace", "", "The namespace of the metric to create the anomaly detection model for.")
	cloudwatch_putAnomalyDetectorCmd.Flags().String("single-metric-anomaly-detector", "", "A single metric anomaly detector to be created.")
	cloudwatch_putAnomalyDetectorCmd.Flags().String("stat", "", "The statistic to use for the metric and the anomaly detection model.")
	cloudwatchCmd.AddCommand(cloudwatch_putAnomalyDetectorCmd)
}
