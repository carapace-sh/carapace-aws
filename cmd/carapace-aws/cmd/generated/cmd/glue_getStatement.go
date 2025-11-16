package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getStatementCmd = &cobra.Command{
	Use:   "get-statement",
	Short: "Retrieves the statement.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getStatementCmd).Standalone()

	glue_getStatementCmd.Flags().String("id", "", "The Id of the statement.")
	glue_getStatementCmd.Flags().String("request-origin", "", "The origin of the request.")
	glue_getStatementCmd.Flags().String("session-id", "", "The Session ID of the statement.")
	glue_getStatementCmd.MarkFlagRequired("id")
	glue_getStatementCmd.MarkFlagRequired("session-id")
	glueCmd.AddCommand(glue_getStatementCmd)
}
