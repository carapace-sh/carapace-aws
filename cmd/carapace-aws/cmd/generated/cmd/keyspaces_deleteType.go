package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_deleteTypeCmd = &cobra.Command{
	Use:   "delete-type",
	Short: "The `DeleteType` operation deletes a user-defined type (UDT).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_deleteTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspaces_deleteTypeCmd).Standalone()

		keyspaces_deleteTypeCmd.Flags().String("keyspace-name", "", "The name of the keyspace of the to be deleted type.")
		keyspaces_deleteTypeCmd.Flags().String("type-name", "", "The name of the type to be deleted.")
		keyspaces_deleteTypeCmd.MarkFlagRequired("keyspace-name")
		keyspaces_deleteTypeCmd.MarkFlagRequired("type-name")
	})
	keyspacesCmd.AddCommand(keyspaces_deleteTypeCmd)
}
