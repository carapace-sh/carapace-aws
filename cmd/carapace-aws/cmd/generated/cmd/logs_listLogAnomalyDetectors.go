package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_listLogAnomalyDetectorsCmd = &cobra.Command{
	Use:   "list-log-anomaly-detectors",
	Short: "Retrieves a list of the log anomaly detectors in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_listLogAnomalyDetectorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_listLogAnomalyDetectorsCmd).Standalone()

		logs_listLogAnomalyDetectorsCmd.Flags().String("filter-log-group-arn", "", "Use this to optionally filter the results to only include anomaly detectors that are associated with the specified log group.")
		logs_listLogAnomalyDetectorsCmd.Flags().String("limit", "", "The maximum number of items to return.")
		logs_listLogAnomalyDetectorsCmd.Flags().String("next-token", "", "")
	})
	logsCmd.AddCommand(logs_listLogAnomalyDetectorsCmd)
}
