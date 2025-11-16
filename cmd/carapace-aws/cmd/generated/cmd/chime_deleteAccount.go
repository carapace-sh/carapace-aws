package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_deleteAccountCmd = &cobra.Command{
	Use:   "delete-account",
	Short: "Deletes the specified Amazon Chime account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_deleteAccountCmd).Standalone()

	chime_deleteAccountCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_deleteAccountCmd.MarkFlagRequired("account-id")
	chimeCmd.AddCommand(chime_deleteAccountCmd)
}
