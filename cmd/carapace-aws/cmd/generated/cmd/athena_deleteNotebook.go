package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_deleteNotebookCmd = &cobra.Command{
	Use:   "delete-notebook",
	Short: "Deletes the specified notebook.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_deleteNotebookCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_deleteNotebookCmd).Standalone()

		athena_deleteNotebookCmd.Flags().String("notebook-id", "", "The ID of the notebook to delete.")
		athena_deleteNotebookCmd.MarkFlagRequired("notebook-id")
	})
	athenaCmd.AddCommand(athena_deleteNotebookCmd)
}
