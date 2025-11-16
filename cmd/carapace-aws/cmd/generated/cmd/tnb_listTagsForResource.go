package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags for AWS TNB resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_listTagsForResourceCmd).Standalone()

		tnb_listTagsForResourceCmd.Flags().String("resource-arn", "", "Resource ARN.")
		tnb_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	tnbCmd.AddCommand(tnb_listTagsForResourceCmd)
}
