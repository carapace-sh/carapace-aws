package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_updateAnomalyCmd = &cobra.Command{
	Use:   "update-anomaly",
	Short: "Use this operation to *suppress* anomaly detection for a specified anomaly or pattern.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_updateAnomalyCmd).Standalone()

	logs_updateAnomalyCmd.Flags().String("anomaly-detector-arn", "", "The ARN of the anomaly detector that this operation is to act on.")
	logs_updateAnomalyCmd.Flags().String("anomaly-id", "", "If you are suppressing or unsuppressing an anomaly, specify its unique ID here.")
	logs_updateAnomalyCmd.Flags().String("baseline", "", "Set this to `true` to prevent CloudWatch Logs from displaying this behavior as an anomaly in the future.")
	logs_updateAnomalyCmd.Flags().String("pattern-id", "", "If you are suppressing or unsuppressing an pattern, specify its unique ID here.")
	logs_updateAnomalyCmd.Flags().String("suppression-period", "", "If you are temporarily suppressing an anomaly or pattern, use this structure to specify how long the suppression is to last.")
	logs_updateAnomalyCmd.Flags().String("suppression-type", "", "Use this to specify whether the suppression to be temporary or infinite.")
	logs_updateAnomalyCmd.MarkFlagRequired("anomaly-detector-arn")
	logsCmd.AddCommand(logs_updateAnomalyCmd)
}
