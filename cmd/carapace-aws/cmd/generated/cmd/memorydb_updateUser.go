package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Changes user password(s) and/or access string.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_updateUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_updateUserCmd).Standalone()

		memorydb_updateUserCmd.Flags().String("access-string", "", "Access permissions string used for this user.")
		memorydb_updateUserCmd.Flags().String("authentication-mode", "", "Denotes the user's authentication properties, such as whether it requires a password to authenticate.")
		memorydb_updateUserCmd.Flags().String("user-name", "", "The name of the user")
		memorydb_updateUserCmd.MarkFlagRequired("user-name")
	})
	memorydbCmd.AddCommand(memorydb_updateUserCmd)
}
