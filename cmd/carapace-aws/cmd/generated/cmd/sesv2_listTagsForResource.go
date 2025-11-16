package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieve a list of the tags (keys and values) that are associated with a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listTagsForResourceCmd).Standalone()

	sesv2_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to retrieve tag information for.")
	sesv2_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	sesv2Cmd.AddCommand(sesv2_listTagsForResourceCmd)
}
