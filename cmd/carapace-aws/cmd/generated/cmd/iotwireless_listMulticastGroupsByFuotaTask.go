package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotwireless_listMulticastGroupsByFuotaTaskCmd = &cobra.Command{
	Use:   "list-multicast-groups-by-fuota-task",
	Short: "List all multicast groups associated with a FUOTA task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotwireless_listMulticastGroupsByFuotaTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotwireless_listMulticastGroupsByFuotaTaskCmd).Standalone()

		iotwireless_listMulticastGroupsByFuotaTaskCmd.Flags().String("id", "", "")
		iotwireless_listMulticastGroupsByFuotaTaskCmd.Flags().String("max-results", "", "")
		iotwireless_listMulticastGroupsByFuotaTaskCmd.Flags().String("next-token", "", "To retrieve the next set of results, the `nextToken` value from a previous response; otherwise **null** to receive the first set of results.")
		iotwireless_listMulticastGroupsByFuotaTaskCmd.MarkFlagRequired("id")
	})
	iotwirelessCmd.AddCommand(iotwireless_listMulticastGroupsByFuotaTaskCmd)
}
