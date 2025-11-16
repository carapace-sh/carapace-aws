package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_updateNotebookCmd = &cobra.Command{
	Use:   "update-notebook",
	Short: "Updates the contents of a Spark notebook.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_updateNotebookCmd).Standalone()

	athena_updateNotebookCmd.Flags().String("client-request-token", "", "A unique case-sensitive string used to ensure the request to create the notebook is idempotent (executes only once).")
	athena_updateNotebookCmd.Flags().String("notebook-id", "", "The ID of the notebook to update.")
	athena_updateNotebookCmd.Flags().String("payload", "", "The updated content for the notebook.")
	athena_updateNotebookCmd.Flags().String("session-id", "", "The active notebook session ID.")
	athena_updateNotebookCmd.Flags().String("type", "", "The notebook content type.")
	athena_updateNotebookCmd.MarkFlagRequired("notebook-id")
	athena_updateNotebookCmd.MarkFlagRequired("payload")
	athena_updateNotebookCmd.MarkFlagRequired("type")
	athenaCmd.AddCommand(athena_updateNotebookCmd)
}
