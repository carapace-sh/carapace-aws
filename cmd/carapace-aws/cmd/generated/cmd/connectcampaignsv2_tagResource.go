package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Tag a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_tagResourceCmd).Standalone()

		connectcampaignsv2_tagResourceCmd.Flags().String("arn", "", "")
		connectcampaignsv2_tagResourceCmd.Flags().String("tags", "", "")
		connectcampaignsv2_tagResourceCmd.MarkFlagRequired("arn")
		connectcampaignsv2_tagResourceCmd.MarkFlagRequired("tags")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_tagResourceCmd)
}
