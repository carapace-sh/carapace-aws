package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add one or more tags (keys and values) to a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_tagResourceCmd).Standalone()

		sesv2_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to add one or more tags to.")
		sesv2_tagResourceCmd.Flags().String("tags", "", "A list of the tags that you want to add to the resource.")
		sesv2_tagResourceCmd.MarkFlagRequired("resource-arn")
		sesv2_tagResourceCmd.MarkFlagRequired("tags")
	})
	sesv2Cmd.AddCommand(sesv2_tagResourceCmd)
}
