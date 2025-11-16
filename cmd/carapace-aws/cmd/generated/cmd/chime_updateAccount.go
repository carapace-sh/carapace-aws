package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_updateAccountCmd = &cobra.Command{
	Use:   "update-account",
	Short: "Updates account details for the specified Amazon Chime account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_updateAccountCmd).Standalone()

	chime_updateAccountCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_updateAccountCmd.Flags().String("default-license", "", "The default license applied when you add users to an Amazon Chime account.")
	chime_updateAccountCmd.Flags().String("name", "", "The new name for the specified Amazon Chime account.")
	chime_updateAccountCmd.MarkFlagRequired("account-id")
	chimeCmd.AddCommand(chime_updateAccountCmd)
}
