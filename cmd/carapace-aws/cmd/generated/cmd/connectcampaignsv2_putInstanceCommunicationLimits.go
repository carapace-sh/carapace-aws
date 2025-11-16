package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcampaignsv2_putInstanceCommunicationLimitsCmd = &cobra.Command{
	Use:   "put-instance-communication-limits",
	Short: "Put the instance communication limits.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcampaignsv2_putInstanceCommunicationLimitsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcampaignsv2_putInstanceCommunicationLimitsCmd).Standalone()

		connectcampaignsv2_putInstanceCommunicationLimitsCmd.Flags().String("communication-limits-config", "", "")
		connectcampaignsv2_putInstanceCommunicationLimitsCmd.Flags().String("connect-instance-id", "", "")
		connectcampaignsv2_putInstanceCommunicationLimitsCmd.MarkFlagRequired("communication-limits-config")
		connectcampaignsv2_putInstanceCommunicationLimitsCmd.MarkFlagRequired("connect-instance-id")
	})
	connectcampaignsv2Cmd.AddCommand(connectcampaignsv2_putInstanceCommunicationLimitsCmd)
}
