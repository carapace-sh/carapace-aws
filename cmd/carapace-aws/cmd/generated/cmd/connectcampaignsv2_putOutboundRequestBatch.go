package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_putOutboundRequestBatchCmd = &cobra.Command{
	Use:   "put-outbound-request-batch",
	Short: "Creates outbound requests for the specified campaign Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_putOutboundRequestBatchCmd).Standalone()

	connectcampaignsv2_putOutboundRequestBatchCmd.Flags().String("id", "", "")
	connectcampaignsv2_putOutboundRequestBatchCmd.Flags().String("outbound-requests", "", "")
	connectcampaignsv2_putOutboundRequestBatchCmd.MarkFlagRequired("id")
	connectcampaignsv2_putOutboundRequestBatchCmd.MarkFlagRequired("outbound-requests")
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_putOutboundRequestBatchCmd)
}
