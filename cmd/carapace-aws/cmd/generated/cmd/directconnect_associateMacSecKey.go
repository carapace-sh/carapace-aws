package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_associateMacSecKeyCmd = &cobra.Command{
	Use:   "associate-mac-sec-key",
	Short: "Associates a MAC Security (MACsec) Connection Key Name (CKN)/ Connectivity Association Key (CAK) pair with a Direct Connect connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_associateMacSecKeyCmd).Standalone()

	directconnect_associateMacSecKeyCmd.Flags().String("cak", "", "The MAC Security (MACsec) CAK to associate with the connection.")
	directconnect_associateMacSecKeyCmd.Flags().String("ckn", "", "The MAC Security (MACsec) CKN to associate with the connection.")
	directconnect_associateMacSecKeyCmd.Flags().String("connection-id", "", "The ID of the dedicated connection (dxcon-xxxx), interconnect (dxcon-xxxx), or LAG (dxlag-xxxx).")
	directconnect_associateMacSecKeyCmd.Flags().String("secret-arn", "", "The Amazon Resource Name (ARN) of the MAC Security (MACsec) secret key to associate with the connection.")
	directconnect_associateMacSecKeyCmd.MarkFlagRequired("connection-id")
	directconnectCmd.AddCommand(directconnect_associateMacSecKeyCmd)
}
