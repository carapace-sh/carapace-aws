package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_cancelQueryCmd = &cobra.Command{
	Use:   "cancel-query",
	Short: "Cancels a query that has been issued.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_cancelQueryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamQuery_cancelQueryCmd).Standalone()

		timestreamQuery_cancelQueryCmd.Flags().String("query-id", "", "The ID of the query that needs to be cancelled.")
		timestreamQuery_cancelQueryCmd.MarkFlagRequired("query-id")
	})
	timestreamQueryCmd.AddCommand(timestreamQuery_cancelQueryCmd)
}
