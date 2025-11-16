package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dlm_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dlm_listTagsForResourceCmd).Standalone()

	dlm_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	dlm_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	dlmCmd.AddCommand(dlm_listTagsForResourceCmd)
}
