package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_removeTagsFromResourceCmd = &cobra.Command{
	Use:   "remove-tags-from-resource",
	Short: "Removes metadata tags from an Amazon RDS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_removeTagsFromResourceCmd).Standalone()

	rds_removeTagsFromResourceCmd.Flags().String("resource-name", "", "The Amazon RDS resource that the tags are removed from.")
	rds_removeTagsFromResourceCmd.Flags().String("tag-keys", "", "The tag key (name) of the tag to be removed.")
	rds_removeTagsFromResourceCmd.MarkFlagRequired("resource-name")
	rds_removeTagsFromResourceCmd.MarkFlagRequired("tag-keys")
	rdsCmd.AddCommand(rds_removeTagsFromResourceCmd)
}
