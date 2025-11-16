package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes the specified set of tags from the specified set of Elastic Disaster Recovery resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_untagResourceCmd).Standalone()

	drs_untagResourceCmd.Flags().String("resource-arn", "", "ARN of the resource for which tags are to be removed.")
	drs_untagResourceCmd.Flags().String("tag-keys", "", "Array of tags to be removed.")
	drs_untagResourceCmd.MarkFlagRequired("resource-arn")
	drs_untagResourceCmd.MarkFlagRequired("tag-keys")
	drsCmd.AddCommand(drs_untagResourceCmd)
}
