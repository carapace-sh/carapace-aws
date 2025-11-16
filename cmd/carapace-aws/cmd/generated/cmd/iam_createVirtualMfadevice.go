package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_createVirtualMfadeviceCmd = &cobra.Command{
	Use:   "create-virtual-mfadevice",
	Short: "Creates a new virtual MFA device for the Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_createVirtualMfadeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_createVirtualMfadeviceCmd).Standalone()

		iam_createVirtualMfadeviceCmd.Flags().String("path", "", "The path for the virtual MFA device.")
		iam_createVirtualMfadeviceCmd.Flags().String("tags", "", "A list of tags that you want to attach to the new IAM virtual MFA device.")
		iam_createVirtualMfadeviceCmd.Flags().String("virtual-mfadevice-name", "", "The name of the virtual MFA device, which must be unique.")
		iam_createVirtualMfadeviceCmd.MarkFlagRequired("virtual-mfadevice-name")
	})
	iamCmd.AddCommand(iam_createVirtualMfadeviceCmd)
}
