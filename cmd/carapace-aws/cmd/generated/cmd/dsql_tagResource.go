package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags a resource with a map of key and value pairs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsql_tagResourceCmd).Standalone()

		dsql_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource that you want to tag.")
		dsql_tagResourceCmd.Flags().String("tags", "", "A map of key and value pairs to use to tag your resource.")
		dsql_tagResourceCmd.MarkFlagRequired("resource-arn")
		dsql_tagResourceCmd.MarkFlagRequired("tags")
	})
	dsqlCmd.AddCommand(dsql_tagResourceCmd)
}
