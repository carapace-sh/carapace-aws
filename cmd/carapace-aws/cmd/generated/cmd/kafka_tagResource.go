package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to the specified MSK resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_tagResourceCmd).Standalone()

	kafka_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the resource that's associated with the tags.")
	kafka_tagResourceCmd.Flags().String("tags", "", "The key-value pair for the resource tag.")
	kafka_tagResourceCmd.MarkFlagRequired("resource-arn")
	kafka_tagResourceCmd.MarkFlagRequired("tags")
	kafkaCmd.AddCommand(kafka_tagResourceCmd)
}
