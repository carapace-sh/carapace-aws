package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iot_putVerificationStateOnViolationCmd = &cobra.Command{
	Use:   "put-verification-state-on-violation",
	Short: "Set a verification state and provide a description of that verification state on a violation (detect alarm).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iot_putVerificationStateOnViolationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iot_putVerificationStateOnViolationCmd).Standalone()

		iot_putVerificationStateOnViolationCmd.Flags().String("verification-state", "", "The verification state of the violation.")
		iot_putVerificationStateOnViolationCmd.Flags().String("verification-state-description", "", "The description of the verification state of the violation (detect alarm).")
		iot_putVerificationStateOnViolationCmd.Flags().String("violation-id", "", "The violation ID.")
		iot_putVerificationStateOnViolationCmd.MarkFlagRequired("verification-state")
		iot_putVerificationStateOnViolationCmd.MarkFlagRequired("violation-id")
	})
	iotCmd.AddCommand(iot_putVerificationStateOnViolationCmd)
}
