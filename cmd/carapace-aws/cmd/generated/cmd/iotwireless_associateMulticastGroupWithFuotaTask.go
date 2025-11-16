package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_associateMulticastGroupWithFuotaTaskCmd = &cobra.Command{
	Use:   "associate-multicast-group-with-fuota-task",
	Short: "Associate a multicast group with a FUOTA task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_associateMulticastGroupWithFuotaTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_associateMulticastGroupWithFuotaTaskCmd).Standalone()

		iotwireless_associateMulticastGroupWithFuotaTaskCmd.Flags().String("id", "", "")
		iotwireless_associateMulticastGroupWithFuotaTaskCmd.Flags().String("multicast-group-id", "", "")
		iotwireless_associateMulticastGroupWithFuotaTaskCmd.MarkFlagRequired("id")
		iotwireless_associateMulticastGroupWithFuotaTaskCmd.MarkFlagRequired("multicast-group-id")
	})
	iotwirelessCmd.AddCommand(iotwireless_associateMulticastGroupWithFuotaTaskCmd)
}
