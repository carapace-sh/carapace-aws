package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_updateNotebookMetadataCmd = &cobra.Command{
	Use:   "update-notebook-metadata",
	Short: "Updates the metadata for a notebook.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_updateNotebookMetadataCmd).Standalone()

	athena_updateNotebookMetadataCmd.Flags().String("client-request-token", "", "A unique case-sensitive string used to ensure the request to create the notebook is idempotent (executes only once).")
	athena_updateNotebookMetadataCmd.Flags().String("name", "", "The name to update the notebook to.")
	athena_updateNotebookMetadataCmd.Flags().String("notebook-id", "", "The ID of the notebook to update the metadata for.")
	athena_updateNotebookMetadataCmd.MarkFlagRequired("name")
	athena_updateNotebookMetadataCmd.MarkFlagRequired("notebook-id")
	athenaCmd.AddCommand(athena_updateNotebookMetadataCmd)
}
