package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_deleteVirtualMfadeviceCmd = &cobra.Command{
	Use:   "delete-virtual-mfadevice",
	Short: "Deletes a virtual MFA device.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_deleteVirtualMfadeviceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iam_deleteVirtualMfadeviceCmd).Standalone()

		iam_deleteVirtualMfadeviceCmd.Flags().String("serial-number", "", "The serial number that uniquely identifies the MFA device.")
		iam_deleteVirtualMfadeviceCmd.MarkFlagRequired("serial-number")
	})
	iamCmd.AddCommand(iam_deleteVirtualMfadeviceCmd)
}
