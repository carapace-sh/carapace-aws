package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshiftData_cancelStatementCmd = &cobra.Command{
	Use:   "cancel-statement",
	Short: "Cancels a running query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshiftData_cancelStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshiftData_cancelStatementCmd).Standalone()

		redshiftData_cancelStatementCmd.Flags().String("id", "", "The identifier of the SQL statement to cancel.")
		redshiftData_cancelStatementCmd.MarkFlagRequired("id")
	})
	redshiftDataCmd.AddCommand(redshiftData_cancelStatementCmd)
}
