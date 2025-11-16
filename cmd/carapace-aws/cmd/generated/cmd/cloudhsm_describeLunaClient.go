package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsm_describeLunaClientCmd = &cobra.Command{
	Use:   "describe-luna-client",
	Short: "This is documentation for **AWS CloudHSM Classic**.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsm_describeLunaClientCmd).Standalone()

	cloudhsm_describeLunaClientCmd.Flags().String("certificate-fingerprint", "", "The certificate fingerprint.")
	cloudhsm_describeLunaClientCmd.Flags().String("client-arn", "", "The ARN of the client.")
	cloudhsmCmd.AddCommand(cloudhsm_describeLunaClientCmd)
}
