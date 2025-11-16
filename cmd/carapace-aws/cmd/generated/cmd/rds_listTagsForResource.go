package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rds_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags on an Amazon RDS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rds_listTagsForResourceCmd).Standalone()

	rds_listTagsForResourceCmd.Flags().String("filters", "", "This parameter isn't currently supported.")
	rds_listTagsForResourceCmd.Flags().String("resource-name", "", "The Amazon RDS resource with tags to be listed.")
	rds_listTagsForResourceCmd.MarkFlagRequired("resource-name")
	rdsCmd.AddCommand(rds_listTagsForResourceCmd)
}
