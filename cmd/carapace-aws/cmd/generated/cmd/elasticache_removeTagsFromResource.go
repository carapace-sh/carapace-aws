package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_removeTagsFromResourceCmd = &cobra.Command{
	Use:   "remove-tags-from-resource",
	Short: "Removes the tags identified by the `TagKeys` list from the named resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_removeTagsFromResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticache_removeTagsFromResourceCmd).Standalone()

		elasticache_removeTagsFromResourceCmd.Flags().String("resource-name", "", "The Amazon Resource Name (ARN) of the resource from which you want the tags removed, for example `arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster` or `arn:aws:elasticache:us-west-2:0123456789:snapshot:mySnapshot`.")
		elasticache_removeTagsFromResourceCmd.Flags().String("tag-keys", "", "A list of `TagKeys` identifying the tags you want removed from the named resource.")
		elasticache_removeTagsFromResourceCmd.MarkFlagRequired("resource-name")
		elasticache_removeTagsFromResourceCmd.MarkFlagRequired("tag-keys")
	})
	elasticacheCmd.AddCommand(elasticache_removeTagsFromResourceCmd)
}
