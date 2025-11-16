package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeUserCmd = &cobra.Command{
	Use:   "describe-user",
	Short: "Describes the user assigned to the specific file transfer protocol-enabled server, as identified by its `ServerId` property.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_describeUserCmd).Standalone()

		transfer_describeUserCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server that has this user assigned.")
		transfer_describeUserCmd.Flags().String("user-name", "", "The name of the user assigned to one or more servers.")
		transfer_describeUserCmd.MarkFlagRequired("server-id")
		transfer_describeUserCmd.MarkFlagRequired("user-name")
	})
	transferCmd.AddCommand(transfer_describeUserCmd)
}
