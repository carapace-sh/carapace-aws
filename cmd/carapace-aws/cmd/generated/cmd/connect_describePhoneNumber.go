package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describePhoneNumberCmd = &cobra.Command{
	Use:   "describe-phone-number",
	Short: "Gets details and status of a phone number that’s claimed to your Amazon Connect instance or traffic distribution group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describePhoneNumberCmd).Standalone()

	connect_describePhoneNumberCmd.Flags().String("phone-number-id", "", "A unique identifier for the phone number.")
	connect_describePhoneNumberCmd.MarkFlagRequired("phone-number-id")
	connectCmd.AddCommand(connect_describePhoneNumberCmd)
}
