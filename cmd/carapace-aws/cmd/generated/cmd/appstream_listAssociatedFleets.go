package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_listAssociatedFleetsCmd = &cobra.Command{
	Use:   "list-associated-fleets",
	Short: "Retrieves the name of the fleet that is associated with the specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_listAssociatedFleetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_listAssociatedFleetsCmd).Standalone()

		appstream_listAssociatedFleetsCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
		appstream_listAssociatedFleetsCmd.Flags().String("stack-name", "", "The name of the stack.")
		appstream_listAssociatedFleetsCmd.MarkFlagRequired("stack-name")
	})
	appstreamCmd.AddCommand(appstream_listAssociatedFleetsCmd)
}
