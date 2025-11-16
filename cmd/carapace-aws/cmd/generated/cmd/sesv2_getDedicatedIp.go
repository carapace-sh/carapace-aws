package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_getDedicatedIpCmd = &cobra.Command{
	Use:   "get-dedicated-ip",
	Short: "Get information about a dedicated IP address, including the name of the dedicated IP pool that it's associated with, as well information about the automatic warm-up process for the address.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_getDedicatedIpCmd).Standalone()

	sesv2_getDedicatedIpCmd.Flags().String("ip", "", "The IP address that you want to obtain more information about.")
	sesv2_getDedicatedIpCmd.MarkFlagRequired("ip")
	sesv2Cmd.AddCommand(sesv2_getDedicatedIpCmd)
}
