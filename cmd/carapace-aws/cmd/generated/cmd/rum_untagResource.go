package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rum_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rum_untagResourceCmd).Standalone()

	rum_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch RUM resource that you're removing tags from.")
	rum_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
	rum_untagResourceCmd.MarkFlagRequired("resource-arn")
	rum_untagResourceCmd.MarkFlagRequired("tag-keys")
	rumCmd.AddCommand(rum_untagResourceCmd)
}
