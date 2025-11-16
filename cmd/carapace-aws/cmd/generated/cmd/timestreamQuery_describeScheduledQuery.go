package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_describeScheduledQueryCmd = &cobra.Command{
	Use:   "describe-scheduled-query",
	Short: "Provides detailed information about a scheduled query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_describeScheduledQueryCmd).Standalone()

	timestreamQuery_describeScheduledQueryCmd.Flags().String("scheduled-query-arn", "", "The ARN of the scheduled query.")
	timestreamQuery_describeScheduledQueryCmd.MarkFlagRequired("scheduled-query-arn")
	timestreamQueryCmd.AddCommand(timestreamQuery_describeScheduledQueryCmd)
}
