package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudwatch_deleteAnomalyDetectorCmd = &cobra.Command{
	Use:   "delete-anomaly-detector",
	Short: "Deletes the specified anomaly detection model from your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudwatch_deleteAnomalyDetectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudwatch_deleteAnomalyDetectorCmd).Standalone()

		cloudwatch_deleteAnomalyDetectorCmd.Flags().String("dimensions", "", "The metric dimensions associated with the anomaly detection model to delete.")
		cloudwatch_deleteAnomalyDetectorCmd.Flags().String("metric-math-anomaly-detector", "", "The metric math anomaly detector to be deleted.")
		cloudwatch_deleteAnomalyDetectorCmd.Flags().String("metric-name", "", "The metric name associated with the anomaly detection model to delete.")
		cloudwatch_deleteAnomalyDetectorCmd.Flags().String("namespace", "", "The namespace associated with the anomaly detection model to delete.")
		cloudwatch_deleteAnomalyDetectorCmd.Flags().String("single-metric-anomaly-detector", "", "A single metric anomaly detector to be deleted.")
		cloudwatch_deleteAnomalyDetectorCmd.Flags().String("stat", "", "The statistic associated with the anomaly detection model to delete.")
	})
	cloudwatchCmd.AddCommand(cloudwatch_deleteAnomalyDetectorCmd)
}
