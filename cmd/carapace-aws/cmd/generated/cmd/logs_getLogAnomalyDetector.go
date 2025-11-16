package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getLogAnomalyDetectorCmd = &cobra.Command{
	Use:   "get-log-anomaly-detector",
	Short: "Retrieves information about the log anomaly detector that you specify.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getLogAnomalyDetectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_getLogAnomalyDetectorCmd).Standalone()

		logs_getLogAnomalyDetectorCmd.Flags().String("anomaly-detector-arn", "", "The ARN of the anomaly detector to retrieve information about.")
		logs_getLogAnomalyDetectorCmd.MarkFlagRequired("anomaly-detector-arn")
	})
	logsCmd.AddCommand(logs_getLogAnomalyDetectorCmd)
}
