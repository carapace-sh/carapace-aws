package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_stopArchiveSearchCmd = &cobra.Command{
	Use:   "stop-archive-search",
	Short: "Stops an in-progress archive search job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_stopArchiveSearchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanager_stopArchiveSearchCmd).Standalone()

		mailmanager_stopArchiveSearchCmd.Flags().String("search-id", "", "The identifier of the search job to stop.")
		mailmanager_stopArchiveSearchCmd.MarkFlagRequired("search-id")
	})
	mailmanagerCmd.AddCommand(mailmanager_stopArchiveSearchCmd)
}
