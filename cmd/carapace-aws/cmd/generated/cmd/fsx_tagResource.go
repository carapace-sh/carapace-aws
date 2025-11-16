package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tags an Amazon FSx resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsx_tagResourceCmd).Standalone()

		fsx_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon FSx resource that you want to tag.")
		fsx_tagResourceCmd.Flags().String("tags", "", "A list of tags for the resource.")
		fsx_tagResourceCmd.MarkFlagRequired("resource-arn")
		fsx_tagResourceCmd.MarkFlagRequired("tags")
	})
	fsxCmd.AddCommand(fsx_tagResourceCmd)
}
