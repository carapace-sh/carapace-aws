package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeServerCmd = &cobra.Command{
	Use:   "describe-server",
	Short: "Describes a file transfer protocol-enabled server that you specify by passing the `ServerId` parameter.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeServerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_describeServerCmd).Standalone()

		transfer_describeServerCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server.")
		transfer_describeServerCmd.MarkFlagRequired("server-id")
	})
	transferCmd.AddCommand(transfer_describeServerCmd)
}
