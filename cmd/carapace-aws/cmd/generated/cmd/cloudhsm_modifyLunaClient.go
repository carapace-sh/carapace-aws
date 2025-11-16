package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_modifyLunaClientCmd = &cobra.Command{
	Use:   "modify-luna-client",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_modifyLunaClientCmd).Standalone()

	cloudhsm_modifyLunaClientCmd.Flags().String("certificate", "", "The new certificate for the client.")
	cloudhsm_modifyLunaClientCmd.Flags().String("client-arn", "", "The ARN of the client.")
	cloudhsm_modifyLunaClientCmd.MarkFlagRequired("certificate")
	cloudhsm_modifyLunaClientCmd.MarkFlagRequired("client-arn")
	cloudhsmCmd.AddCommand(cloudhsm_modifyLunaClientCmd)
}
