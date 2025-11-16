package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_putAnomalyDetectorCmd = &cobra.Command{
	Use:   "put-anomaly-detector",
	Short: "When you call `PutAnomalyDetector`, the operation creates a new anomaly detector if one doesn't exist, or updates an existing one.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_putAnomalyDetectorCmd).Standalone()

	amp_putAnomalyDetectorCmd.Flags().String("anomaly-detector-id", "", "The identifier of the anomaly detector to update.")
	amp_putAnomalyDetectorCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	amp_putAnomalyDetectorCmd.Flags().String("configuration", "", "The algorithm configuration for the anomaly detector.")
	amp_putAnomalyDetectorCmd.Flags().String("evaluation-interval-in-seconds", "", "The frequency, in seconds, at which the anomaly detector evaluates metrics.")
	amp_putAnomalyDetectorCmd.Flags().String("labels", "", "The Amazon Managed Service for Prometheus metric labels to associate with the anomaly detector.")
	amp_putAnomalyDetectorCmd.Flags().String("missing-data-action", "", "Specifies the action to take when data is missing during evaluation.")
	amp_putAnomalyDetectorCmd.Flags().String("workspace-id", "", "The identifier of the workspace containing the anomaly detector to update.")
	amp_putAnomalyDetectorCmd.MarkFlagRequired("anomaly-detector-id")
	amp_putAnomalyDetectorCmd.MarkFlagRequired("configuration")
	amp_putAnomalyDetectorCmd.MarkFlagRequired("workspace-id")
	ampCmd.AddCommand(amp_putAnomalyDetectorCmd)
}
