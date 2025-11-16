package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_addTagsToResourceCmd = &cobra.Command{
	Use:   "add-tags-to-resource",
	Short: "A tag is a key-value pair where the key and value are case-sensitive.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_addTagsToResourceCmd).Standalone()

	elasticache_addTagsToResourceCmd.Flags().String("resource-name", "", "The Amazon Resource Name (ARN) of the resource to which the tags are to be added, for example `arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster` or `arn:aws:elasticache:us-west-2:0123456789:snapshot:mySnapshot`.")
	elasticache_addTagsToResourceCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	elasticache_addTagsToResourceCmd.MarkFlagRequired("resource-name")
	elasticache_addTagsToResourceCmd.MarkFlagRequired("tags")
	elasticacheCmd.AddCommand(elasticache_addTagsToResourceCmd)
}
