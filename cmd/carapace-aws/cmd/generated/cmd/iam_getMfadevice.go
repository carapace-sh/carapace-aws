package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iam_getMfadeviceCmd = &cobra.Command{
	Use:   "get-mfadevice",
	Short: "Retrieves information about an MFA device for a specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iam_getMfadeviceCmd).Standalone()

	iam_getMfadeviceCmd.Flags().String("serial-number", "", "Serial number that uniquely identifies the MFA device.")
	iam_getMfadeviceCmd.Flags().String("user-name", "", "The friendly name identifying the user.")
	iam_getMfadeviceCmd.MarkFlagRequired("serial-number")
	iamCmd.AddCommand(iam_getMfadeviceCmd)
}
