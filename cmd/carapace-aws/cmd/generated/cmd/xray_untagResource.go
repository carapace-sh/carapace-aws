package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from an Amazon Web Services X-Ray group or sampling rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_untagResourceCmd).Standalone()

	xray_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Number (ARN) of an X-Ray group or sampling rule.")
	xray_untagResourceCmd.Flags().String("tag-keys", "", "Keys for one or more tags that you want to remove from an X-Ray group or sampling rule.")
	xray_untagResourceCmd.MarkFlagRequired("resource-arn")
	xray_untagResourceCmd.MarkFlagRequired("tag-keys")
	xrayCmd.AddCommand(xray_untagResourceCmd)
}
