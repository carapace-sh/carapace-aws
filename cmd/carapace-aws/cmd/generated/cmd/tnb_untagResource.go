package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tnb_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untags an AWS TNB resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tnb_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(tnb_untagResourceCmd).Standalone()

		tnb_untagResourceCmd.Flags().String("resource-arn", "", "Resource ARN.")
		tnb_untagResourceCmd.Flags().String("tag-keys", "", "Tag keys.")
		tnb_untagResourceCmd.MarkFlagRequired("resource-arn")
		tnb_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	tnbCmd.AddCommand(tnb_untagResourceCmd)
}
