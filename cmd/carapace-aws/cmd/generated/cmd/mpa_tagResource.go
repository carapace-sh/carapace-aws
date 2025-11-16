package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpa_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Creates or updates a resource tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpa_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpa_tagResourceCmd).Standalone()

		mpa_tagResourceCmd.Flags().String("resource-arn", "", "Amazon Resource Name (ARN) for the resource you want to tag.")
		mpa_tagResourceCmd.Flags().String("tags", "", "Tags that you have added to the specified resource.")
		mpa_tagResourceCmd.MarkFlagRequired("resource-arn")
		mpa_tagResourceCmd.MarkFlagRequired("tags")
	})
	mpaCmd.AddCommand(mpa_tagResourceCmd)
}
