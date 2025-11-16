package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_listLogGroupsForQueryCmd = &cobra.Command{
	Use:   "list-log-groups-for-query",
	Short: "Returns a list of the log groups that were analyzed during a single CloudWatch Logs Insights query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_listLogGroupsForQueryCmd).Standalone()

	logs_listLogGroupsForQueryCmd.Flags().String("max-results", "", "Limits the number of returned log groups to the specified number.")
	logs_listLogGroupsForQueryCmd.Flags().String("next-token", "", "")
	logs_listLogGroupsForQueryCmd.Flags().String("query-id", "", "The ID of the query to use.")
	logs_listLogGroupsForQueryCmd.MarkFlagRequired("query-id")
	logsCmd.AddCommand(logs_listLogGroupsForQueryCmd)
}
