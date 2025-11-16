package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_putProfileOutboundRequestBatchCmd = &cobra.Command{
	Use:   "put-profile-outbound-request-batch",
	Short: "Takes in a list of profile outbound requests to be placed as part of an outbound campaign.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_putProfileOutboundRequestBatchCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_putProfileOutboundRequestBatchCmd).Standalone()

		connectcampaignsv2_putProfileOutboundRequestBatchCmd.Flags().String("id", "", "")
		connectcampaignsv2_putProfileOutboundRequestBatchCmd.Flags().String("profile-outbound-requests", "", "")
		connectcampaignsv2_putProfileOutboundRequestBatchCmd.MarkFlagRequired("id")
		connectcampaignsv2_putProfileOutboundRequestBatchCmd.MarkFlagRequired("profile-outbound-requests")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_putProfileOutboundRequestBatchCmd)
}
