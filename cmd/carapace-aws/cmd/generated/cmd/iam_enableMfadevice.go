package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_enableMfadeviceCmd = &cobra.Command{
	Use:   "enable-mfadevice",
	Short: "Enables the specified MFA device and associates it with the specified IAM user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_enableMfadeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_enableMfadeviceCmd).Standalone()

		iam_enableMfadeviceCmd.Flags().String("authentication-code1", "", "An authentication code emitted by the device.")
		iam_enableMfadeviceCmd.Flags().String("authentication-code2", "", "A subsequent authentication code emitted by the device.")
		iam_enableMfadeviceCmd.Flags().String("serial-number", "", "The serial number that uniquely identifies the MFA device.")
		iam_enableMfadeviceCmd.Flags().String("user-name", "", "The name of the IAM user for whom you want to enable the MFA device.")
		iam_enableMfadeviceCmd.MarkFlagRequired("authentication-code1")
		iam_enableMfadeviceCmd.MarkFlagRequired("authentication-code2")
		iam_enableMfadeviceCmd.MarkFlagRequired("serial-number")
		iam_enableMfadeviceCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_enableMfadeviceCmd)
}
