package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_listStatementsCmd = &cobra.Command{
	Use:   "list-statements",
	Short: "Lists statements for the session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_listStatementsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(glue_listStatementsCmd).Standalone()

		glue_listStatementsCmd.Flags().String("next-token", "", "A continuation token, if this is a continuation call.")
		glue_listStatementsCmd.Flags().String("request-origin", "", "The origin of the request to list statements.")
		glue_listStatementsCmd.Flags().String("session-id", "", "The Session ID of the statements.")
		glue_listStatementsCmd.MarkFlagRequired("session-id")
	})
	glueCmd.AddCommand(glue_listStatementsCmd)
}
