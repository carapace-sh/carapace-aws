package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDataExports_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List tags associated with an existing data export.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDataExports_listTagsForResourceCmd).Standalone()

	bcmDataExports_listTagsForResourceCmd.Flags().String("max-results", "", "The maximum number of objects that are returned for the request.")
	bcmDataExports_listTagsForResourceCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	bcmDataExports_listTagsForResourceCmd.Flags().String("resource-arn", "", "The unique identifier for the resource.")
	bcmDataExports_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	bcmDataExportsCmd.AddCommand(bcmDataExports_listTagsForResourceCmd)
}
