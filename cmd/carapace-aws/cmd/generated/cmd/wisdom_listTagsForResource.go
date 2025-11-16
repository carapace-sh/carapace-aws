package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wisdom_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wisdom_listTagsForResourceCmd).Standalone()

	wisdom_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	wisdom_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	wisdomCmd.AddCommand(wisdom_listTagsForResourceCmd)
}
