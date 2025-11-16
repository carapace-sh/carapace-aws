package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_deleteLogAnomalyDetectorCmd = &cobra.Command{
	Use:   "delete-log-anomaly-detector",
	Short: "Deletes the specified CloudWatch Logs anomaly detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_deleteLogAnomalyDetectorCmd).Standalone()

	logs_deleteLogAnomalyDetectorCmd.Flags().String("anomaly-detector-arn", "", "The ARN of the anomaly detector to delete.")
	logs_deleteLogAnomalyDetectorCmd.MarkFlagRequired("anomaly-detector-arn")
	logsCmd.AddCommand(logs_deleteLogAnomalyDetectorCmd)
}
