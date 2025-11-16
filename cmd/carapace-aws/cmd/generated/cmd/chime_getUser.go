package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_getUserCmd = &cobra.Command{
	Use:   "get-user",
	Short: "Retrieves details for the specified user ID, such as primary email address, license type,and personal meeting PIN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_getUserCmd).Standalone()

	chime_getUserCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_getUserCmd.Flags().String("user-id", "", "The user ID.")
	chime_getUserCmd.MarkFlagRequired("account-id")
	chime_getUserCmd.MarkFlagRequired("user-id")
	chimeCmd.AddCommand(chime_getUserCmd)
}
