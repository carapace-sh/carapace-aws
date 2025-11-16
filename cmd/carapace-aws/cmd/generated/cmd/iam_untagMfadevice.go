package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_untagMfadeviceCmd = &cobra.Command{
	Use:   "untag-mfadevice",
	Short: "Removes the specified tags from the IAM virtual multi-factor authentication (MFA) device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_untagMfadeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_untagMfadeviceCmd).Standalone()

		iam_untagMfadeviceCmd.Flags().String("serial-number", "", "The unique identifier for the IAM virtual MFA device from which you want to remove tags.")
		iam_untagMfadeviceCmd.Flags().String("tag-keys", "", "A list of key names as a simple array of strings.")
		iam_untagMfadeviceCmd.MarkFlagRequired("serial-number")
		iam_untagMfadeviceCmd.MarkFlagRequired("tag-keys")
	})
	iamCmd.AddCommand(iam_untagMfadeviceCmd)
}
