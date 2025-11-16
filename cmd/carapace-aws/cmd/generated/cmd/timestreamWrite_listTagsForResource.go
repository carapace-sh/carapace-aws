package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags on a Timestream resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_listTagsForResourceCmd).Standalone()

	timestreamWrite_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Timestream resource with tags to be listed.")
	timestreamWrite_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	timestreamWriteCmd.AddCommand(timestreamWrite_listTagsForResourceCmd)
}
