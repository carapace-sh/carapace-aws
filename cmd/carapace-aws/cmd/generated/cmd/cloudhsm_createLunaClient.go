package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_createLunaClientCmd = &cobra.Command{
	Use:   "create-luna-client",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_createLunaClientCmd).Standalone()

	cloudhsm_createLunaClientCmd.Flags().String("certificate", "", "The contents of a Base64-Encoded X.509 v3 certificate to be installed on the HSMs used by this client.")
	cloudhsm_createLunaClientCmd.Flags().String("label", "", "The label for the client.")
	cloudhsm_createLunaClientCmd.MarkFlagRequired("certificate")
	cloudhsmCmd.AddCommand(cloudhsm_createLunaClientCmd)
}
