package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_getNotebookMetadataCmd = &cobra.Command{
	Use:   "get-notebook-metadata",
	Short: "Retrieves notebook metadata for the specified notebook ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_getNotebookMetadataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_getNotebookMetadataCmd).Standalone()

		athena_getNotebookMetadataCmd.Flags().String("notebook-id", "", "The ID of the notebook whose metadata is to be retrieved.")
		athena_getNotebookMetadataCmd.MarkFlagRequired("notebook-id")
	})
	athenaCmd.AddCommand(athena_getNotebookMetadataCmd)
}
