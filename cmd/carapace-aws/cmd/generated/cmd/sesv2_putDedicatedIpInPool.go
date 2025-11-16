package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putDedicatedIpInPoolCmd = &cobra.Command{
	Use:   "put-dedicated-ip-in-pool",
	Short: "Move a dedicated IP address to an existing dedicated IP pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putDedicatedIpInPoolCmd).Standalone()

	sesv2_putDedicatedIpInPoolCmd.Flags().String("destination-pool-name", "", "The name of the IP pool that you want to add the dedicated IP address to.")
	sesv2_putDedicatedIpInPoolCmd.Flags().String("ip", "", "The IP address that you want to move to the dedicated IP pool.")
	sesv2_putDedicatedIpInPoolCmd.MarkFlagRequired("destination-pool-name")
	sesv2_putDedicatedIpInPoolCmd.MarkFlagRequired("ip")
	sesv2Cmd.AddCommand(sesv2_putDedicatedIpInPoolCmd)
}
