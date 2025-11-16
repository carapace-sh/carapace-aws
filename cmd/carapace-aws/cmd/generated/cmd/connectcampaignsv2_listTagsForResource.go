package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_listTagsForResourceCmd).Standalone()

		connectcampaignsv2_listTagsForResourceCmd.Flags().String("arn", "", "")
		connectcampaignsv2_listTagsForResourceCmd.MarkFlagRequired("arn")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_listTagsForResourceCmd)
}
