package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_tagMfadeviceCmd = &cobra.Command{
	Use:   "tag-mfadevice",
	Short: "Adds one or more tags to an IAM virtual multi-factor authentication (MFA) device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_tagMfadeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_tagMfadeviceCmd).Standalone()

		iam_tagMfadeviceCmd.Flags().String("serial-number", "", "The unique identifier for the IAM virtual MFA device to which you want to add tags.")
		iam_tagMfadeviceCmd.Flags().String("tags", "", "The list of tags that you want to attach to the IAM virtual MFA device.")
		iam_tagMfadeviceCmd.MarkFlagRequired("serial-number")
		iam_tagMfadeviceCmd.MarkFlagRequired("tags")
	})
	iamCmd.AddCommand(iam_tagMfadeviceCmd)
}
