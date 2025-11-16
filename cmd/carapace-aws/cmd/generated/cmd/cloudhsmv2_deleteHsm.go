package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_deleteHsmCmd = &cobra.Command{
	Use:   "delete-hsm",
	Short: "Deletes the specified HSM.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_deleteHsmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudhsmv2_deleteHsmCmd).Standalone()

		cloudhsmv2_deleteHsmCmd.Flags().String("cluster-id", "", "The identifier (ID) of the cluster that contains the HSM that you are deleting.")
		cloudhsmv2_deleteHsmCmd.Flags().String("eni-id", "", "The identifier (ID) of the elastic network interface (ENI) of the HSM that you are deleting.")
		cloudhsmv2_deleteHsmCmd.Flags().String("eni-ip", "", "The IP address of the elastic network interface (ENI) of the HSM that you are deleting.")
		cloudhsmv2_deleteHsmCmd.Flags().String("hsm-id", "", "The identifier (ID) of the HSM that you are deleting.")
		cloudhsmv2_deleteHsmCmd.MarkFlagRequired("cluster-id")
	})
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_deleteHsmCmd)
}
