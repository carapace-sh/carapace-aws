package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Untag a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_untagResourceCmd).Standalone()

	connectcampaigns_untagResourceCmd.Flags().String("arn", "", "")
	connectcampaigns_untagResourceCmd.Flags().String("tag-keys", "", "")
	connectcampaigns_untagResourceCmd.MarkFlagRequired("arn")
	connectcampaigns_untagResourceCmd.MarkFlagRequired("tag-keys")
	connectcampaignsCmd.AddCommand(connectcampaigns_untagResourceCmd)
}
