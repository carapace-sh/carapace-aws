package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_tagLogGroupCmd = &cobra.Command{
	Use:   "tag-log-group",
	Short: "The TagLogGroup operation is on the path to deprecation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_tagLogGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_tagLogGroupCmd).Standalone()

		logs_tagLogGroupCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_tagLogGroupCmd.Flags().String("tags", "", "The key-value pairs to use for the tags.")
		logs_tagLogGroupCmd.MarkFlagRequired("log-group-name")
		logs_tagLogGroupCmd.MarkFlagRequired("tags")
	})
	logsCmd.AddCommand(logs_tagLogGroupCmd)
}
