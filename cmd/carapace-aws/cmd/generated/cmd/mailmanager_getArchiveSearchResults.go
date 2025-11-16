package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanager_getArchiveSearchResultsCmd = &cobra.Command{
	Use:   "get-archive-search-results",
	Short: "Returns the results of a completed email archive search job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanager_getArchiveSearchResultsCmd).Standalone()

	mailmanager_getArchiveSearchResultsCmd.Flags().String("search-id", "", "The identifier of the completed search job.")
	mailmanager_getArchiveSearchResultsCmd.MarkFlagRequired("search-id")
	mailmanagerCmd.AddCommand(mailmanager_getArchiveSearchResultsCmd)
}
