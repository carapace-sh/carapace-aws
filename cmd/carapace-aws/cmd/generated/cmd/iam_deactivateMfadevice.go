package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deactivateMfadeviceCmd = &cobra.Command{
	Use:   "deactivate-mfadevice",
	Short: "Deactivates the specified MFA device and removes it from association with the user name for which it was originally enabled.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deactivateMfadeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_deactivateMfadeviceCmd).Standalone()

		iam_deactivateMfadeviceCmd.Flags().String("serial-number", "", "The serial number that uniquely identifies the MFA device.")
		iam_deactivateMfadeviceCmd.Flags().String("user-name", "", "The name of the user whose MFA device you want to deactivate.")
		iam_deactivateMfadeviceCmd.MarkFlagRequired("serial-number")
	})
	iamCmd.AddCommand(iam_deactivateMfadeviceCmd)
}
