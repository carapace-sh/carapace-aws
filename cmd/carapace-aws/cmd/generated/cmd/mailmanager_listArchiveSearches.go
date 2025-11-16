package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listArchiveSearchesCmd = &cobra.Command{
	Use:   "list-archive-searches",
	Short: "Returns a list of email archive search jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listArchiveSearchesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_listArchiveSearchesCmd).Standalone()

		mailmanager_listArchiveSearchesCmd.Flags().String("archive-id", "", "The identifier of the archive.")
		mailmanager_listArchiveSearchesCmd.Flags().String("next-token", "", "If NextToken is returned, there are more results available.")
		mailmanager_listArchiveSearchesCmd.Flags().String("page-size", "", "The maximum number of archive search jobs that are returned per call.")
		mailmanager_listArchiveSearchesCmd.MarkFlagRequired("archive-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_listArchiveSearchesCmd)
}
