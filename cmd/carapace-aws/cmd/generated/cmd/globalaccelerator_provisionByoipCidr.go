package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_provisionByoipCidrCmd = &cobra.Command{
	Use:   "provision-byoip-cidr",
	Short: "Provisions an IP address range to use with your Amazon Web Services resources through bring your own IP addresses (BYOIP) and creates a corresponding address pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_provisionByoipCidrCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_provisionByoipCidrCmd).Standalone()

		globalaccelerator_provisionByoipCidrCmd.Flags().String("cidr", "", "The public IPv4 address range, in CIDR notation.")
		globalaccelerator_provisionByoipCidrCmd.Flags().String("cidr-authorization-context", "", "A signed document that proves that you are authorized to bring the specified IP address range to Amazon using BYOIP.")
		globalaccelerator_provisionByoipCidrCmd.MarkFlagRequired("cidr")
		globalaccelerator_provisionByoipCidrCmd.MarkFlagRequired("cidr-authorization-context")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_provisionByoipCidrCmd)
}
