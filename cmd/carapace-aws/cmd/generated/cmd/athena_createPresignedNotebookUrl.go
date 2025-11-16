package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_createPresignedNotebookUrlCmd = &cobra.Command{
	Use:   "create-presigned-notebook-url",
	Short: "Gets an authentication token and the URL at which the notebook can be accessed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_createPresignedNotebookUrlCmd).Standalone()

	athena_createPresignedNotebookUrlCmd.Flags().String("session-id", "", "The session ID.")
	athena_createPresignedNotebookUrlCmd.MarkFlagRequired("session-id")
	athenaCmd.AddCommand(athena_createPresignedNotebookUrlCmd)
}
