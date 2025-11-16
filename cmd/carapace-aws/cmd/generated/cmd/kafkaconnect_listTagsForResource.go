package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all the tags attached to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafkaconnect_listTagsForResourceCmd).Standalone()

		kafkaconnect_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which you want to list all attached tags.")
		kafkaconnect_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	kafkaconnectCmd.AddCommand(kafkaconnect_listTagsForResourceCmd)
}
