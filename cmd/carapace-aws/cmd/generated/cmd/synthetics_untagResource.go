package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var synthetics_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(synthetics_untagResourceCmd).Standalone()

	synthetics_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the canary or group that you're removing tags from.")
	synthetics_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
	synthetics_untagResourceCmd.MarkFlagRequired("resource-arn")
	synthetics_untagResourceCmd.MarkFlagRequired("tag-keys")
	syntheticsCmd.AddCommand(synthetics_untagResourceCmd)
}
