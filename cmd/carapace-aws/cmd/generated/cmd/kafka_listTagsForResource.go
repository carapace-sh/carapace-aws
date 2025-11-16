package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of the tags associated with the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_listTagsForResourceCmd).Standalone()

		kafka_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the resource that's associated with the tags.")
		kafka_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	kafkaCmd.AddCommand(kafka_listTagsForResourceCmd)
}
