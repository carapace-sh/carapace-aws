package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_disassociateMulticastGroupFromFuotaTaskCmd = &cobra.Command{
	Use:   "disassociate-multicast-group-from-fuota-task",
	Short: "Disassociates a multicast group from a FUOTA task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_disassociateMulticastGroupFromFuotaTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_disassociateMulticastGroupFromFuotaTaskCmd).Standalone()

		iotwireless_disassociateMulticastGroupFromFuotaTaskCmd.Flags().String("id", "", "")
		iotwireless_disassociateMulticastGroupFromFuotaTaskCmd.Flags().String("multicast-group-id", "", "")
		iotwireless_disassociateMulticastGroupFromFuotaTaskCmd.MarkFlagRequired("id")
		iotwireless_disassociateMulticastGroupFromFuotaTaskCmd.MarkFlagRequired("multicast-group-id")
	})
	iotwirelessCmd.AddCommand(iotwireless_disassociateMulticastGroupFromFuotaTaskCmd)
}
