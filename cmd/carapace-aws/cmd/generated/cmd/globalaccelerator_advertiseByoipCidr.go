package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_advertiseByoipCidrCmd = &cobra.Command{
	Use:   "advertise-byoip-cidr",
	Short: "Advertises an IPv4 address range that is provisioned for use with your Amazon Web Services resources through bring your own IP addresses (BYOIP).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_advertiseByoipCidrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_advertiseByoipCidrCmd).Standalone()

		globalaccelerator_advertiseByoipCidrCmd.Flags().String("cidr", "", "The address range, in CIDR notation.")
		globalaccelerator_advertiseByoipCidrCmd.MarkFlagRequired("cidr")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_advertiseByoipCidrCmd)
}
