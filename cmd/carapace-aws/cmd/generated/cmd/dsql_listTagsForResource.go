package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all of the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_listTagsForResourceCmd).Standalone()

	dsql_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource for which you want to list the tags.")
	dsql_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	dsqlCmd.AddCommand(dsql_listTagsForResourceCmd)
}
