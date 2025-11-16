package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafkaconnect_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Attaches tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafkaconnect_tagResourceCmd).Standalone()

	kafkaconnect_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to which you want to attach tags.")
	kafkaconnect_tagResourceCmd.Flags().String("tags", "", "The tags that you want to attach to the resource.")
	kafkaconnect_tagResourceCmd.MarkFlagRequired("resource-arn")
	kafkaconnect_tagResourceCmd.MarkFlagRequired("tags")
	kafkaconnectCmd.AddCommand(kafkaconnect_tagResourceCmd)
}
