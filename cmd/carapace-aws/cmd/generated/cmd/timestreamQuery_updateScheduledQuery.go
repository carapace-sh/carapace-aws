package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_updateScheduledQueryCmd = &cobra.Command{
	Use:   "update-scheduled-query",
	Short: "Update a scheduled query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_updateScheduledQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamQuery_updateScheduledQueryCmd).Standalone()

		timestreamQuery_updateScheduledQueryCmd.Flags().String("scheduled-query-arn", "", "ARN of the scheuled query.")
		timestreamQuery_updateScheduledQueryCmd.Flags().String("state", "", "State of the scheduled query.")
		timestreamQuery_updateScheduledQueryCmd.MarkFlagRequired("scheduled-query-arn")
		timestreamQuery_updateScheduledQueryCmd.MarkFlagRequired("state")
	})
	timestreamQueryCmd.AddCommand(timestreamQuery_updateScheduledQueryCmd)
}
