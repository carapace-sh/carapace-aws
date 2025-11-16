package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsql_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsql_untagResourceCmd).Standalone()

	dsql_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource from which to remove tags.")
	dsql_untagResourceCmd.Flags().String("tag-keys", "", "The array of keys of the tags that you want to remove.")
	dsql_untagResourceCmd.MarkFlagRequired("resource-arn")
	dsql_untagResourceCmd.MarkFlagRequired("tag-keys")
	dsqlCmd.AddCommand(dsql_untagResourceCmd)
}
