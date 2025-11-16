package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var athena_exportNotebookCmd = &cobra.Command{
	Use:   "export-notebook",
	Short: "Exports the specified notebook and its metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(athena_exportNotebookCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(athena_exportNotebookCmd).Standalone()

		athena_exportNotebookCmd.Flags().String("notebook-id", "", "The ID of the notebook to export.")
		athena_exportNotebookCmd.MarkFlagRequired("notebook-id")
	})
	athenaCmd.AddCommand(athena_exportNotebookCmd)
}
