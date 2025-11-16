package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_resetPersonalPinCmd = &cobra.Command{
	Use:   "reset-personal-pin",
	Short: "Resets the personal meeting PIN for the specified user on an Amazon Chime account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_resetPersonalPinCmd).Standalone()

	chime_resetPersonalPinCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_resetPersonalPinCmd.Flags().String("user-id", "", "The user ID.")
	chime_resetPersonalPinCmd.MarkFlagRequired("account-id")
	chime_resetPersonalPinCmd.MarkFlagRequired("user-id")
	chimeCmd.AddCommand(chime_resetPersonalPinCmd)
}
