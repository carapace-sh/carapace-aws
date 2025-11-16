package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes a user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_deleteUserCmd).Standalone()

	memorydb_deleteUserCmd.Flags().String("user-name", "", "The name of the user to delete")
	memorydb_deleteUserCmd.MarkFlagRequired("user-name")
	memorydbCmd.AddCommand(memorydb_deleteUserCmd)
}
