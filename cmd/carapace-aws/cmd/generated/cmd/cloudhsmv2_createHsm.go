package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_createHsmCmd = &cobra.Command{
	Use:   "create-hsm",
	Short: "Creates a new hardware security module (HSM) in the specified CloudHSM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_createHsmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsmv2_createHsmCmd).Standalone()

		cloudhsmv2_createHsmCmd.Flags().String("availability-zone", "", "The Availability Zone where you are creating the HSM.")
		cloudhsmv2_createHsmCmd.Flags().String("cluster-id", "", "The identifier (ID) of the HSM's cluster.")
		cloudhsmv2_createHsmCmd.Flags().String("ip-address", "", "The HSM's IP address.")
		cloudhsmv2_createHsmCmd.MarkFlagRequired("availability-zone")
		cloudhsmv2_createHsmCmd.MarkFlagRequired("cluster-id")
	})
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_createHsmCmd)
}
