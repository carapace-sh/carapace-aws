package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_cancelStatementCmd = &cobra.Command{
	Use:   "cancel-statement",
	Short: "Cancels the statement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_cancelStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_cancelStatementCmd).Standalone()

		glue_cancelStatementCmd.Flags().String("id", "", "The ID of the statement to be cancelled.")
		glue_cancelStatementCmd.Flags().String("request-origin", "", "The origin of the request to cancel the statement.")
		glue_cancelStatementCmd.Flags().String("session-id", "", "The Session ID of the statement to be cancelled.")
		glue_cancelStatementCmd.MarkFlagRequired("id")
		glue_cancelStatementCmd.MarkFlagRequired("session-id")
	})
	glueCmd.AddCommand(glue_cancelStatementCmd)
}
