package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a user under the specified Amazon Chime account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_createUserCmd).Standalone()

	chime_createUserCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_createUserCmd.Flags().String("email", "", "The user's email address.")
	chime_createUserCmd.Flags().String("user-type", "", "The user type.")
	chime_createUserCmd.Flags().String("username", "", "The user name.")
	chime_createUserCmd.MarkFlagRequired("account-id")
	chimeCmd.AddCommand(chime_createUserCmd)
}
