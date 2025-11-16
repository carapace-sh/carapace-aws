package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transcribe_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more custom tags, each in the form of a key:value pair, to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transcribe_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transcribe_tagResourceCmd).Standalone()

		transcribe_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource you want to tag.")
		transcribe_tagResourceCmd.Flags().String("tags", "", "Adds one or more custom tags, each in the form of a key:value pair, to the specified resource.")
		transcribe_tagResourceCmd.MarkFlagRequired("resource-arn")
		transcribe_tagResourceCmd.MarkFlagRequired("tags")
	})
	transcribeCmd.AddCommand(transcribe_tagResourceCmd)
}
