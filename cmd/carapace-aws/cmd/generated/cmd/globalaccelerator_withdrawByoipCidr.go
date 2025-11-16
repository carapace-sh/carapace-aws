package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_withdrawByoipCidrCmd = &cobra.Command{
	Use:   "withdraw-byoip-cidr",
	Short: "Stops advertising an address range that is provisioned as an address pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_withdrawByoipCidrCmd).Standalone()

	globalaccelerator_withdrawByoipCidrCmd.Flags().String("cidr", "", "The address range, in CIDR notation.")
	globalaccelerator_withdrawByoipCidrCmd.MarkFlagRequired("cidr")
	globalacceleratorCmd.AddCommand(globalaccelerator_withdrawByoipCidrCmd)
}
