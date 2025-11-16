package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getArchiveSearchCmd = &cobra.Command{
	Use:   "get-archive-search",
	Short: "Retrieves the details and current status of a specific email archive search job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getArchiveSearchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_getArchiveSearchCmd).Standalone()

		mailmanager_getArchiveSearchCmd.Flags().String("search-id", "", "The identifier of the search job to get details for.")
		mailmanager_getArchiveSearchCmd.MarkFlagRequired("search-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_getArchiveSearchCmd)
}
