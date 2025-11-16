package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeHostKeyCmd = &cobra.Command{
	Use:   "describe-host-key",
	Short: "Returns the details of the host key that's specified by the `HostKeyId` and `ServerId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeHostKeyCmd).Standalone()

	transfer_describeHostKeyCmd.Flags().String("host-key-id", "", "The identifier of the host key that you want described.")
	transfer_describeHostKeyCmd.Flags().String("server-id", "", "The identifier of the server that contains the host key that you want described.")
	transfer_describeHostKeyCmd.MarkFlagRequired("host-key-id")
	transfer_describeHostKeyCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_describeHostKeyCmd)
}
