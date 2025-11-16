package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_untagLogGroupCmd = &cobra.Command{
	Use:   "untag-log-group",
	Short: "The UntagLogGroup operation is on the path to deprecation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_untagLogGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_untagLogGroupCmd).Standalone()

		logs_untagLogGroupCmd.Flags().String("log-group-name", "", "The name of the log group.")
		logs_untagLogGroupCmd.Flags().String("tags", "", "The tag keys.")
		logs_untagLogGroupCmd.MarkFlagRequired("log-group-name")
		logs_untagLogGroupCmd.MarkFlagRequired("tags")
	})
	logsCmd.AddCommand(logs_untagLogGroupCmd)
}
