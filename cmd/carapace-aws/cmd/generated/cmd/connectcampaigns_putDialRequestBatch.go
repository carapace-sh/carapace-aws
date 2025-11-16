package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaigns_putDialRequestBatchCmd = &cobra.Command{
	Use:   "put-dial-request-batch",
	Short: "Creates dials requests for the specified campaign Amazon Connect account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaigns_putDialRequestBatchCmd).Standalone()

	connectcampaigns_putDialRequestBatchCmd.Flags().String("dial-requests", "", "")
	connectcampaigns_putDialRequestBatchCmd.Flags().String("id", "", "")
	connectcampaigns_putDialRequestBatchCmd.MarkFlagRequired("dial-requests")
	connectcampaigns_putDialRequestBatchCmd.MarkFlagRequired("id")
	connectcampaignsCmd.AddCommand(connectcampaigns_putDialRequestBatchCmd)
}
