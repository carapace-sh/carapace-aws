package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_untagResourceCmd).Standalone()

	evidently_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch Evidently resource that you're removing tags from.")
	evidently_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
	evidently_untagResourceCmd.MarkFlagRequired("resource-arn")
	evidently_untagResourceCmd.MarkFlagRequired("tag-keys")
	evidentlyCmd.AddCommand(evidently_untagResourceCmd)
}
