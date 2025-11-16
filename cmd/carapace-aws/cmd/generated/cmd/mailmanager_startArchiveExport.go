package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_startArchiveExportCmd = &cobra.Command{
	Use:   "start-archive-export",
	Short: "Initiates an export of emails from the specified archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_startArchiveExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_startArchiveExportCmd).Standalone()

		mailmanager_startArchiveExportCmd.Flags().String("archive-id", "", "The identifier of the archive to export emails from.")
		mailmanager_startArchiveExportCmd.Flags().String("export-destination-configuration", "", "Details on where to deliver the exported email data.")
		mailmanager_startArchiveExportCmd.Flags().String("filters", "", "Criteria to filter which emails are included in the export.")
		mailmanager_startArchiveExportCmd.Flags().String("from-timestamp", "", "The start of the timestamp range to include emails from.")
		mailmanager_startArchiveExportCmd.Flags().Bool("include-metadata", false, "Whether to include message metadata as JSON files in the export.")
		mailmanager_startArchiveExportCmd.Flags().String("max-results", "", "The maximum number of email items to include in the export.")
		mailmanager_startArchiveExportCmd.Flags().Bool("no-include-metadata", false, "Whether to include message metadata as JSON files in the export.")
		mailmanager_startArchiveExportCmd.Flags().String("to-timestamp", "", "The end of the timestamp range to include emails from.")
		mailmanager_startArchiveExportCmd.MarkFlagRequired("archive-id")
		mailmanager_startArchiveExportCmd.MarkFlagRequired("export-destination-configuration")
		mailmanager_startArchiveExportCmd.MarkFlagRequired("from-timestamp")
		mailmanager_startArchiveExportCmd.Flag("no-include-metadata").Hidden = true
		mailmanager_startArchiveExportCmd.MarkFlagRequired("to-timestamp")
	})
	mailmanagerCmd.AddCommand(mailmanager_startArchiveExportCmd)
}
