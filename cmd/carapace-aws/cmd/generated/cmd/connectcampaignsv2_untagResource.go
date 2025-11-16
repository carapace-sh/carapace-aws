package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untag a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_untagResourceCmd).Standalone()

		connectcampaignsv2_untagResourceCmd.Flags().String("arn", "", "")
		connectcampaignsv2_untagResourceCmd.Flags().String("tag-keys", "", "")
		connectcampaignsv2_untagResourceCmd.MarkFlagRequired("arn")
		connectcampaignsv2_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_untagResourceCmd)
}
