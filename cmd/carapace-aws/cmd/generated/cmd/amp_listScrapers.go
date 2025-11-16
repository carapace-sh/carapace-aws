package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_listScrapersCmd = &cobra.Command{
	Use:   "list-scrapers",
	Short: "The `ListScrapers` operation lists all of the scrapers in your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_listScrapersCmd).Standalone()

	amp_listScrapersCmd.Flags().String("filters", "", "(Optional) A list of key-value pairs to filter the list of scrapers returned.")
	amp_listScrapersCmd.Flags().String("max-results", "", "Optional) The maximum number of scrapers to return in one `ListScrapers` operation.")
	amp_listScrapersCmd.Flags().String("next-token", "", "(Optional) The token for the next set of items to return.")
	ampCmd.AddCommand(amp_listScrapersCmd)
}
