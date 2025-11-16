package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of tags that are applied to the specified Amazon Web Services X-Ray group or sampling rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_listTagsForResourceCmd).Standalone()

		xray_listTagsForResourceCmd.Flags().String("next-token", "", "A pagination token.")
		xray_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Number (ARN) of an X-Ray group or sampling rule.")
		xray_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	xrayCmd.AddCommand(xray_listTagsForResourceCmd)
}
