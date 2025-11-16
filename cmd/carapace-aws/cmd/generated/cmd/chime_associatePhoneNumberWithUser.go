package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_associatePhoneNumberWithUserCmd = &cobra.Command{
	Use:   "associate-phone-number-with-user",
	Short: "Associates a phone number with the specified Amazon Chime user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_associatePhoneNumberWithUserCmd).Standalone()

	chime_associatePhoneNumberWithUserCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_associatePhoneNumberWithUserCmd.Flags().String("e164-phone-number", "", "The phone number, in E.164 format.")
	chime_associatePhoneNumberWithUserCmd.Flags().String("user-id", "", "The user ID.")
	chime_associatePhoneNumberWithUserCmd.MarkFlagRequired("account-id")
	chime_associatePhoneNumberWithUserCmd.MarkFlagRequired("e164-phone-number")
	chime_associatePhoneNumberWithUserCmd.MarkFlagRequired("user-id")
	chimeCmd.AddCommand(chime_associatePhoneNumberWithUserCmd)
}
