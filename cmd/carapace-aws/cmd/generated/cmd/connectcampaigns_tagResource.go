package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tag a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaigns_tagResourceCmd).Standalone()

		connectcampaigns_tagResourceCmd.Flags().String("arn", "", "")
		connectcampaigns_tagResourceCmd.Flags().String("tags", "", "")
		connectcampaigns_tagResourceCmd.MarkFlagRequired("arn")
		connectcampaigns_tagResourceCmd.MarkFlagRequired("tags")
	})
	connectcampaignsCmd.AddCommand(connectcampaigns_tagResourceCmd)
}
