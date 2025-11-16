package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_listTagsLogGroupCmd = &cobra.Command{
	Use:   "list-tags-log-group",
	Short: "The ListTagsLogGroup operation is on the path to deprecation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_listTagsLogGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_listTagsLogGroupCmd).Standalone()

		logs_listTagsLogGroupCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_listTagsLogGroupCmd.MarkFlagRequired("log-group-name")
	})
	logsCmd.AddCommand(logs_listTagsLogGroupCmd)
}
