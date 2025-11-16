package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_resyncMfadeviceCmd = &cobra.Command{
	Use:   "resync-mfadevice",
	Short: "Synchronizes the specified MFA device with its IAM resource object on the Amazon Web Services servers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_resyncMfadeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_resyncMfadeviceCmd).Standalone()

		iam_resyncMfadeviceCmd.Flags().String("authentication-code1", "", "An authentication code emitted by the device.")
		iam_resyncMfadeviceCmd.Flags().String("authentication-code2", "", "A subsequent authentication code emitted by the device.")
		iam_resyncMfadeviceCmd.Flags().String("serial-number", "", "Serial number that uniquely identifies the MFA device.")
		iam_resyncMfadeviceCmd.Flags().String("user-name", "", "The name of the user whose MFA device you want to resynchronize.")
		iam_resyncMfadeviceCmd.MarkFlagRequired("authentication-code1")
		iam_resyncMfadeviceCmd.MarkFlagRequired("authentication-code2")
		iam_resyncMfadeviceCmd.MarkFlagRequired("serial-number")
		iam_resyncMfadeviceCmd.MarkFlagRequired("user-name")
	})
	iamCmd.AddCommand(iam_resyncMfadeviceCmd)
}
