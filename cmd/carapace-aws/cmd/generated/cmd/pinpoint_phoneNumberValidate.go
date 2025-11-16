package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpoint_phoneNumberValidateCmd = &cobra.Command{
	Use:   "phone-number-validate",
	Short: "Retrieves information about a phone number.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpoint_phoneNumberValidateCmd).Standalone()

	pinpoint_phoneNumberValidateCmd.Flags().String("number-validate-request", "", "")
	pinpoint_phoneNumberValidateCmd.MarkFlagRequired("number-validate-request")
	pinpointCmd.AddCommand(pinpoint_phoneNumberValidateCmd)
}
