package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_logoutUserCmd = &cobra.Command{
	Use:   "logout-user",
	Short: "Logs out the specified user from all of the devices they are currently logged into.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_logoutUserCmd).Standalone()

	chime_logoutUserCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_logoutUserCmd.Flags().String("user-id", "", "The user ID.")
	chime_logoutUserCmd.MarkFlagRequired("account-id")
	chime_logoutUserCmd.MarkFlagRequired("user-id")
	chimeCmd.AddCommand(chime_logoutUserCmd)
}
