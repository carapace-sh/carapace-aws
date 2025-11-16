package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backupsearch_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "This operation puts tags on the resource you indicate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backupsearch_tagResourceCmd).Standalone()

	backupsearch_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the resource.")
	backupsearch_tagResourceCmd.Flags().String("tags", "", "Required tags to include.")
	backupsearch_tagResourceCmd.MarkFlagRequired("resource-arn")
	backupsearch_tagResourceCmd.MarkFlagRequired("tags")
	backupsearchCmd.AddCommand(backupsearch_tagResourceCmd)
}
