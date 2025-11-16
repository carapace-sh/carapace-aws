package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_deprovisionByoipCidrCmd = &cobra.Command{
	Use:   "deprovision-byoip-cidr",
	Short: "Releases the specified address range that you provisioned to use with your Amazon Web Services resources through bring your own IP addresses (BYOIP) and deletes the corresponding address pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_deprovisionByoipCidrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_deprovisionByoipCidrCmd).Standalone()

		globalaccelerator_deprovisionByoipCidrCmd.Flags().String("cidr", "", "The address range, in CIDR notation.")
		globalaccelerator_deprovisionByoipCidrCmd.MarkFlagRequired("cidr")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_deprovisionByoipCidrCmd)
}
