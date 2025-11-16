package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var timestreamQuery_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the association of tags from a Timestream query resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(timestreamQuery_untagResourceCmd).Standalone()

	timestreamQuery_untagResourceCmd.Flags().String("resource-arn", "", "The Timestream resource that the tags will be removed from.")
	timestreamQuery_untagResourceCmd.Flags().String("tag-keys", "", "A list of tags keys.")
	timestreamQuery_untagResourceCmd.MarkFlagRequired("resource-arn")
	timestreamQuery_untagResourceCmd.MarkFlagRequired("tag-keys")
	timestreamQueryCmd.AddCommand(timestreamQuery_untagResourceCmd)
}
