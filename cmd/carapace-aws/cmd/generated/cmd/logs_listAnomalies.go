package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_listAnomaliesCmd = &cobra.Command{
	Use:   "list-anomalies",
	Short: "Returns a list of anomalies that log anomaly detectors have found.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_listAnomaliesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_listAnomaliesCmd).Standalone()

		logs_listAnomaliesCmd.Flags().String("anomaly-detector-arn", "", "Use this to optionally limit the results to only the anomalies found by a certain anomaly detector.")
		logs_listAnomaliesCmd.Flags().String("limit", "", "The maximum number of items to return.")
		logs_listAnomaliesCmd.Flags().String("next-token", "", "")
		logs_listAnomaliesCmd.Flags().String("suppression-state", "", "You can specify this parameter if you want to the operation to return only anomalies that are currently either suppressed or unsuppressed.")
	})
	logsCmd.AddCommand(logs_listAnomaliesCmd)
}
