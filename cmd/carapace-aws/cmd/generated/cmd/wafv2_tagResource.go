package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associates tags with the specified Amazon Web Services resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_tagResourceCmd).Standalone()

	wafv2_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	wafv2_tagResourceCmd.Flags().String("tags", "", "An array of key:value pairs to associate with the resource.")
	wafv2_tagResourceCmd.MarkFlagRequired("resource-arn")
	wafv2_tagResourceCmd.MarkFlagRequired("tags")
	wafv2Cmd.AddCommand(wafv2_tagResourceCmd)
}
