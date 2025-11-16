package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_listKeysCmd = &cobra.Command{
	Use:   "list-keys",
	Short: "Lists API key resources in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_listKeysCmd).Standalone()

	location_listKeysCmd.Flags().String("filter", "", "Optionally filter the list to only `Active` or `Expired` API keys.")
	location_listKeysCmd.Flags().String("max-results", "", "An optional limit for the number of resources returned in a single call.")
	location_listKeysCmd.Flags().String("next-token", "", "The pagination token specifying which page of results to return in the response.")
	locationCmd.AddCommand(location_listKeysCmd)
}
