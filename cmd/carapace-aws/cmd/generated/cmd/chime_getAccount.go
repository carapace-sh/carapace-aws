package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_getAccountCmd = &cobra.Command{
	Use:   "get-account",
	Short: "Retrieves details for the specified Amazon Chime account, such as account type and supported licenses.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_getAccountCmd).Standalone()

	chime_getAccountCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_getAccountCmd.MarkFlagRequired("account-id")
	chimeCmd.AddCommand(chime_getAccountCmd)
}
