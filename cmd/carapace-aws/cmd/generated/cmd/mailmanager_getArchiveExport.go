package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getArchiveExportCmd = &cobra.Command{
	Use:   "get-archive-export",
	Short: "Retrieves the details and current status of a specific email archive export job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getArchiveExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_getArchiveExportCmd).Standalone()

		mailmanager_getArchiveExportCmd.Flags().String("export-id", "", "The identifier of the export job to get details for.")
		mailmanager_getArchiveExportCmd.MarkFlagRequired("export-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_getArchiveExportCmd)
}
