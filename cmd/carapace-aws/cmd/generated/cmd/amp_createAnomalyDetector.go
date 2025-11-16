package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_createAnomalyDetectorCmd = &cobra.Command{
	Use:   "create-anomaly-detector",
	Short: "Creates an anomaly detector within a workspace using the Random Cut Forest algorithm for time-series analysis.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_createAnomalyDetectorCmd).Standalone()

	amp_createAnomalyDetectorCmd.Flags().String("alias", "", "A user-friendly name for the anomaly detector.")
	amp_createAnomalyDetectorCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	amp_createAnomalyDetectorCmd.Flags().String("configuration", "", "The algorithm configuration for the anomaly detector.")
	amp_createAnomalyDetectorCmd.Flags().String("evaluation-interval-in-seconds", "", "The frequency, in seconds, at which the anomaly detector evaluates metrics.")
	amp_createAnomalyDetectorCmd.Flags().String("labels", "", "The Amazon Managed Service for Prometheus metric labels to associate with the anomaly detector.")
	amp_createAnomalyDetectorCmd.Flags().String("missing-data-action", "", "Specifies the action to take when data is missing during evaluation.")
	amp_createAnomalyDetectorCmd.Flags().String("tags", "", "The metadata to apply to the anomaly detector to assist with categorization and organization.")
	amp_createAnomalyDetectorCmd.Flags().String("workspace-id", "", "The identifier of the workspace where the anomaly detector will be created.")
	amp_createAnomalyDetectorCmd.MarkFlagRequired("alias")
	amp_createAnomalyDetectorCmd.MarkFlagRequired("configuration")
	amp_createAnomalyDetectorCmd.MarkFlagRequired("workspace-id")
	ampCmd.AddCommand(amp_createAnomalyDetectorCmd)
}
