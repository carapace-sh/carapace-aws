package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticache_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags currently on a named resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticache_listTagsForResourceCmd).Standalone()

	elasticache_listTagsForResourceCmd.Flags().String("resource-name", "", "The Amazon Resource Name (ARN) of the resource for which you want the list of tags, for example `arn:aws:elasticache:us-west-2:0123456789:cluster:myCluster` or `arn:aws:elasticache:us-west-2:0123456789:snapshot:mySnapshot`.")
	elasticache_listTagsForResourceCmd.MarkFlagRequired("resource-name")
	elasticacheCmd.AddCommand(elasticache_listTagsForResourceCmd)
}
