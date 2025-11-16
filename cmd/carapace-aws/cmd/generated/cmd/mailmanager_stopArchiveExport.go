package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_stopArchiveExportCmd = &cobra.Command{
	Use:   "stop-archive-export",
	Short: "Stops an in-progress export of emails from an archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_stopArchiveExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_stopArchiveExportCmd).Standalone()

		mailmanager_stopArchiveExportCmd.Flags().String("export-id", "", "The identifier of the export job to stop.")
		mailmanager_stopArchiveExportCmd.MarkFlagRequired("export-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_stopArchiveExportCmd)
}
