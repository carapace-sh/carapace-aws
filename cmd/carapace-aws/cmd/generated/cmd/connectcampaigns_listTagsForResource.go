package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaigns_listTagsForResourceCmd).Standalone()

		connectcampaigns_listTagsForResourceCmd.Flags().String("arn", "", "")
		connectcampaigns_listTagsForResourceCmd.MarkFlagRequired("arn")
	})
	connectcampaignsCmd.AddCommand(connectcampaigns_listTagsForResourceCmd)
}
