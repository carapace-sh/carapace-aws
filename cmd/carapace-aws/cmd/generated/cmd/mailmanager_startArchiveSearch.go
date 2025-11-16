package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_startArchiveSearchCmd = &cobra.Command{
	Use:   "start-archive-search",
	Short: "Initiates a search across emails in the specified archive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_startArchiveSearchCmd).Standalone()

	mailmanager_startArchiveSearchCmd.Flags().String("archive-id", "", "The identifier of the archive to search emails in.")
	mailmanager_startArchiveSearchCmd.Flags().String("filters", "", "Criteria to filter which emails are included in the search results.")
	mailmanager_startArchiveSearchCmd.Flags().String("from-timestamp", "", "The start timestamp of the range to search emails from.")
	mailmanager_startArchiveSearchCmd.Flags().String("max-results", "", "The maximum number of search results to return.")
	mailmanager_startArchiveSearchCmd.Flags().String("to-timestamp", "", "The end timestamp of the range to search emails from.")
	mailmanager_startArchiveSearchCmd.MarkFlagRequired("archive-id")
	mailmanager_startArchiveSearchCmd.MarkFlagRequired("from-timestamp")
	mailmanager_startArchiveSearchCmd.MarkFlagRequired("max-results")
	mailmanager_startArchiveSearchCmd.MarkFlagRequired("to-timestamp")
	mailmanagerCmd.AddCommand(mailmanager_startArchiveSearchCmd)
}
