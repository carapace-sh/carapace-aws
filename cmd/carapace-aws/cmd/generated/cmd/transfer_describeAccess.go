package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeAccessCmd = &cobra.Command{
	Use:   "describe-access",
	Short: "Describes the access that is assigned to the specific file transfer protocol-enabled server, as identified by its `ServerId` property and its `ExternalId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeAccessCmd).Standalone()

	transfer_describeAccessCmd.Flags().String("external-id", "", "A unique identifier that is required to identify specific groups within your directory.")
	transfer_describeAccessCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server that has this access assigned.")
	transfer_describeAccessCmd.MarkFlagRequired("external-id")
	transfer_describeAccessCmd.MarkFlagRequired("server-id")
	transferCmd.AddCommand(transfer_describeAccessCmd)
}
