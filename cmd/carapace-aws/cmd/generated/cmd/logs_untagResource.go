package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_untagResourceCmd).Standalone()

		logs_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the CloudWatch Logs resource that you're removing tags from.")
		logs_untagResourceCmd.Flags().String("tag-keys", "", "The list of tag keys to remove from the resource.")
		logs_untagResourceCmd.MarkFlagRequired("resource-arn")
		logs_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	logsCmd.AddCommand(logs_untagResourceCmd)
}
