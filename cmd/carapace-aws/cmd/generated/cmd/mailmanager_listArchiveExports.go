package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_listArchiveExportsCmd = &cobra.Command{
	Use:   "list-archive-exports",
	Short: "Returns a list of email archive export jobs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_listArchiveExportsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_listArchiveExportsCmd).Standalone()

		mailmanager_listArchiveExportsCmd.Flags().String("archive-id", "", "The identifier of the archive.")
		mailmanager_listArchiveExportsCmd.Flags().String("next-token", "", "If NextToken is returned, there are more results available.")
		mailmanager_listArchiveExportsCmd.Flags().String("page-size", "", "The maximum number of archive export jobs that are returned per call.")
		mailmanager_listArchiveExportsCmd.MarkFlagRequired("archive-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_listArchiveExportsCmd)
}
