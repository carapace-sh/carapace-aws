package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeLogStreamsCmd = &cobra.Command{
	Use:   "describe-log-streams",
	Short: "Lists the log streams for the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeLogStreamsCmd).Standalone()

	logs_describeLogStreamsCmd.Flags().String("descending", "", "If the value is true, results are returned in descending order.")
	logs_describeLogStreamsCmd.Flags().String("limit", "", "The maximum number of items returned.")
	logs_describeLogStreamsCmd.Flags().String("log-group-identifier", "", "Specify either the name or ARN of the log group to view.")
	logs_describeLogStreamsCmd.Flags().String("log-group-name", "", "The name of the log group.")
	logs_describeLogStreamsCmd.Flags().String("log-stream-name-prefix", "", "The prefix to match.")
	logs_describeLogStreamsCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	logs_describeLogStreamsCmd.Flags().String("order-by", "", "If the value is `LogStreamName`, the results are ordered by log stream name.")
	logsCmd.AddCommand(logs_describeLogStreamsCmd)
}
