package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanrooms_listConfiguredTableAssociationsCmd = &cobra.Command{
	Use:   "list-configured-table-associations",
	Short: "Lists configured table associations for a membership.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanrooms_listConfiguredTableAssociationsCmd).Standalone()

	cleanrooms_listConfiguredTableAssociationsCmd.Flags().String("max-results", "", "The maximum number of results that are returned for an API request call.")
	cleanrooms_listConfiguredTableAssociationsCmd.Flags().String("membership-identifier", "", "A unique identifier for the membership to list configured table associations for.")
	cleanrooms_listConfiguredTableAssociationsCmd.Flags().String("next-token", "", "The pagination token that's used to fetch the next set of results.")
	cleanrooms_listConfiguredTableAssociationsCmd.MarkFlagRequired("membership-identifier")
	cleanroomsCmd.AddCommand(cleanrooms_listConfiguredTableAssociationsCmd)
}
