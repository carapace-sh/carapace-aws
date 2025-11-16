package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_listAssociatedStacksCmd = &cobra.Command{
	Use:   "list-associated-stacks",
	Short: "Retrieves the name of the stack with which the specified fleet is associated.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_listAssociatedStacksCmd).Standalone()

	appstream_listAssociatedStacksCmd.Flags().String("fleet-name", "", "The name of the fleet.")
	appstream_listAssociatedStacksCmd.Flags().String("next-token", "", "The pagination token to use to retrieve the next page of results for this operation.")
	appstream_listAssociatedStacksCmd.MarkFlagRequired("fleet-name")
	appstreamCmd.AddCommand(appstream_listAssociatedStacksCmd)
}
