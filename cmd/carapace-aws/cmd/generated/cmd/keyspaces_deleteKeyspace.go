package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_deleteKeyspaceCmd = &cobra.Command{
	Use:   "delete-keyspace",
	Short: "The `DeleteKeyspace` operation deletes a keyspace and all of its tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_deleteKeyspaceCmd).Standalone()

	keyspaces_deleteKeyspaceCmd.Flags().String("keyspace-name", "", "The name of the keyspace to be deleted.")
	keyspaces_deleteKeyspaceCmd.MarkFlagRequired("keyspace-name")
	keyspacesCmd.AddCommand(keyspaces_deleteKeyspaceCmd)
}
