package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a MemoryDB user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_createUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(memorydb_createUserCmd).Standalone()

		memorydb_createUserCmd.Flags().String("access-string", "", "Access permissions string used for this user.")
		memorydb_createUserCmd.Flags().String("authentication-mode", "", "Denotes the user's authentication properties, such as whether it requires a password to authenticate.")
		memorydb_createUserCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
		memorydb_createUserCmd.Flags().String("user-name", "", "The name of the user.")
		memorydb_createUserCmd.MarkFlagRequired("access-string")
		memorydb_createUserCmd.MarkFlagRequired("authentication-mode")
		memorydb_createUserCmd.MarkFlagRequired("user-name")
	})
	memorydbCmd.AddCommand(memorydb_createUserCmd)
}
