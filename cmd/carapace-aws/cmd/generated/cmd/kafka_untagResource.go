package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the tags associated with the keys that are provided in the query.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_untagResourceCmd).Standalone()

		kafka_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the resource that's associated with the tags.")
		kafka_untagResourceCmd.Flags().String("tag-keys", "", "Tag keys must be unique for a given cluster.")
		kafka_untagResourceCmd.MarkFlagRequired("resource-arn")
		kafka_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	kafkaCmd.AddCommand(kafka_untagResourceCmd)
}
