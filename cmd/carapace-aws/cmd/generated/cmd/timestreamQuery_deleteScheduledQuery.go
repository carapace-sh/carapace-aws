package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_deleteScheduledQueryCmd = &cobra.Command{
	Use:   "delete-scheduled-query",
	Short: "Deletes a given scheduled query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_deleteScheduledQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamQuery_deleteScheduledQueryCmd).Standalone()

		timestreamQuery_deleteScheduledQueryCmd.Flags().String("scheduled-query-arn", "", "The ARN of the scheduled query.")
		timestreamQuery_deleteScheduledQueryCmd.MarkFlagRequired("scheduled-query-arn")
	})
	timestreamQueryCmd.AddCommand(timestreamQuery_deleteScheduledQueryCmd)
}
