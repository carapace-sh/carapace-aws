package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_getInstanceCommunicationLimitsCmd = &cobra.Command{
	Use:   "get-instance-communication-limits",
	Short: "Get the instance communication limits.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_getInstanceCommunicationLimitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_getInstanceCommunicationLimitsCmd).Standalone()

		connectcampaignsv2_getInstanceCommunicationLimitsCmd.Flags().String("connect-instance-id", "", "")
		connectcampaignsv2_getInstanceCommunicationLimitsCmd.MarkFlagRequired("connect-instance-id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_getInstanceCommunicationLimitsCmd)
}
