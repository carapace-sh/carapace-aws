package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_searchResourceTagsCmd = &cobra.Command{
	Use:   "search-resource-tags",
	Short: "Searches tags used in an Amazon Connect instance using optional search criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_searchResourceTagsCmd).Standalone()

	connect_searchResourceTagsCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_searchResourceTagsCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
	connect_searchResourceTagsCmd.Flags().String("next-token", "", "The token for the next set of results.")
	connect_searchResourceTagsCmd.Flags().String("resource-types", "", "The list of resource types to be used to search tags from.")
	connect_searchResourceTagsCmd.Flags().String("search-criteria", "", "The search criteria to be used to return tags.")
	connect_searchResourceTagsCmd.MarkFlagRequired("instance-id")
	connectCmd.AddCommand(connect_searchResourceTagsCmd)
}
