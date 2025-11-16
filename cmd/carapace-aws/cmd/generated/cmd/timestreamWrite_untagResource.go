package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamWrite_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the association of tags from a Timestream resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamWrite_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(timestreamWrite_untagResourceCmd).Standalone()

		timestreamWrite_untagResourceCmd.Flags().String("resource-arn", "", "The Timestream resource that the tags will be removed from.")
		timestreamWrite_untagResourceCmd.Flags().String("tag-keys", "", "A list of tags keys.")
		timestreamWrite_untagResourceCmd.MarkFlagRequired("resource-arn")
		timestreamWrite_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	timestreamWriteCmd.AddCommand(timestreamWrite_untagResourceCmd)
}
