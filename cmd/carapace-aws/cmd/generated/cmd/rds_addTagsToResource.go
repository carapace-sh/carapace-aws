package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_addTagsToResourceCmd = &cobra.Command{
	Use:   "add-tags-to-resource",
	Short: "Adds metadata tags to an Amazon RDS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_addTagsToResourceCmd).Standalone()

	rds_addTagsToResourceCmd.Flags().String("resource-name", "", "The Amazon RDS resource that the tags are added to.")
	rds_addTagsToResourceCmd.Flags().String("tags", "", "The tags to be assigned to the Amazon RDS resource.")
	rds_addTagsToResourceCmd.MarkFlagRequired("resource-name")
	rds_addTagsToResourceCmd.MarkFlagRequired("tags")
	rdsCmd.AddCommand(rds_addTagsToResourceCmd)
}
