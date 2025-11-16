package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_runStatementCmd = &cobra.Command{
	Use:   "run-statement",
	Short: "Executes the statement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_runStatementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_runStatementCmd).Standalone()

		glue_runStatementCmd.Flags().String("code", "", "The statement code to be run.")
		glue_runStatementCmd.Flags().String("request-origin", "", "The origin of the request.")
		glue_runStatementCmd.Flags().String("session-id", "", "The Session Id of the statement to be run.")
		glue_runStatementCmd.MarkFlagRequired("code")
		glue_runStatementCmd.MarkFlagRequired("session-id")
	})
	glueCmd.AddCommand(glue_runStatementCmd)
}
