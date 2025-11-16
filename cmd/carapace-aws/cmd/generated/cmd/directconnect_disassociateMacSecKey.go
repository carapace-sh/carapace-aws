package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_disassociateMacSecKeyCmd = &cobra.Command{
	Use:   "disassociate-mac-sec-key",
	Short: "Removes the association between a MAC Security (MACsec) security key and a Direct Connect connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_disassociateMacSecKeyCmd).Standalone()

	directconnect_disassociateMacSecKeyCmd.Flags().String("connection-id", "", "The ID of the dedicated connection (dxcon-xxxx), interconnect (dxcon-xxxx), or LAG (dxlag-xxxx).")
	directconnect_disassociateMacSecKeyCmd.Flags().String("secret-arn", "", "The Amazon Resource Name (ARN) of the MAC Security (MACsec) secret key.")
	directconnect_disassociateMacSecKeyCmd.MarkFlagRequired("connection-id")
	directconnect_disassociateMacSecKeyCmd.MarkFlagRequired("secret-arn")
	directconnectCmd.AddCommand(directconnect_disassociateMacSecKeyCmd)
}
