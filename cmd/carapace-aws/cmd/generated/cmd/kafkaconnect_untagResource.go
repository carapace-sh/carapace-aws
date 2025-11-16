package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafkaconnect_untagResourceCmd).Standalone()

		kafkaconnect_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource from which you want to remove tags.")
		kafkaconnect_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags that you want to remove from the resource.")
		kafkaconnect_untagResourceCmd.MarkFlagRequired("resource-arn")
		kafkaconnect_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	kafkaconnectCmd.AddCommand(kafkaconnect_untagResourceCmd)
}
