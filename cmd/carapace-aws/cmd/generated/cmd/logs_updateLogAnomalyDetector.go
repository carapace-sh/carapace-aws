package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_updateLogAnomalyDetectorCmd = &cobra.Command{
	Use:   "update-log-anomaly-detector",
	Short: "Updates an existing log anomaly detector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_updateLogAnomalyDetectorCmd).Standalone()

	logs_updateLogAnomalyDetectorCmd.Flags().String("anomaly-detector-arn", "", "The ARN of the anomaly detector that you want to update.")
	logs_updateLogAnomalyDetectorCmd.Flags().String("anomaly-visibility-time", "", "The number of days to use as the life cycle of anomalies.")
	logs_updateLogAnomalyDetectorCmd.Flags().Bool("enabled", false, "Use this parameter to pause or restart the anomaly detector.")
	logs_updateLogAnomalyDetectorCmd.Flags().String("evaluation-frequency", "", "Specifies how often the anomaly detector runs and look for anomalies.")
	logs_updateLogAnomalyDetectorCmd.Flags().String("filter-pattern", "", "")
	logs_updateLogAnomalyDetectorCmd.Flags().Bool("no-enabled", false, "Use this parameter to pause or restart the anomaly detector.")
	logs_updateLogAnomalyDetectorCmd.MarkFlagRequired("anomaly-detector-arn")
	logs_updateLogAnomalyDetectorCmd.MarkFlagRequired("enabled")
	logs_updateLogAnomalyDetectorCmd.Flag("no-enabled").Hidden = true
	logs_updateLogAnomalyDetectorCmd.MarkFlagRequired("no-enabled")
	logsCmd.AddCommand(logs_updateLogAnomalyDetectorCmd)
}
