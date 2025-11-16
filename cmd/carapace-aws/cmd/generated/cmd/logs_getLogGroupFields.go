package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_getLogGroupFieldsCmd = &cobra.Command{
	Use:   "get-log-group-fields",
	Short: "Returns a list of the fields that are included in log events in the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_getLogGroupFieldsCmd).Standalone()

	logs_getLogGroupFieldsCmd.Flags().String("log-group-identifier", "", "Specify either the name or ARN of the log group to view.")
	logs_getLogGroupFieldsCmd.Flags().String("log-group-name", "", "The name of the log group to search.")
	logs_getLogGroupFieldsCmd.Flags().String("time", "", "The time to set as the center of the query.")
	logsCmd.AddCommand(logs_getLogGroupFieldsCmd)
}
