package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_createLogAnomalyDetectorCmd = &cobra.Command{
	Use:   "create-log-anomaly-detector",
	Short: "Creates an *anomaly detector* that regularly scans one or more log groups and look for patterns and anomalies in the logs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_createLogAnomalyDetectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_createLogAnomalyDetectorCmd).Standalone()

		logs_createLogAnomalyDetectorCmd.Flags().String("anomaly-visibility-time", "", "The number of days to have visibility on an anomaly.")
		logs_createLogAnomalyDetectorCmd.Flags().String("detector-name", "", "A name for this anomaly detector.")
		logs_createLogAnomalyDetectorCmd.Flags().String("evaluation-frequency", "", "Specifies how often the anomaly detector is to run and look for anomalies.")
		logs_createLogAnomalyDetectorCmd.Flags().String("filter-pattern", "", "You can use this parameter to limit the anomaly detection model to examine only log events that match the pattern you specify here.")
		logs_createLogAnomalyDetectorCmd.Flags().String("kms-key-id", "", "Optionally assigns a KMS key to secure this anomaly detector and its findings.")
		logs_createLogAnomalyDetectorCmd.Flags().String("log-group-arn-list", "", "An array containing the ARN of the log group that this anomaly detector will watch.")
		logs_createLogAnomalyDetectorCmd.Flags().String("tags", "", "An optional list of key-value pairs to associate with the resource.")
		logs_createLogAnomalyDetectorCmd.MarkFlagRequired("log-group-arn-list")
	})
	logsCmd.AddCommand(logs_createLogAnomalyDetectorCmd)
}
