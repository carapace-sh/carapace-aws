package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_deleteTableCmd = &cobra.Command{
	Use:   "delete-table",
	Short: "The `DeleteTable` operation deletes a table and all of its data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_deleteTableCmd).Standalone()

	keyspaces_deleteTableCmd.Flags().String("keyspace-name", "", "The name of the keyspace of the to be deleted table.")
	keyspaces_deleteTableCmd.Flags().String("table-name", "", "The name of the table to be deleted.")
	keyspaces_deleteTableCmd.MarkFlagRequired("keyspace-name")
	keyspaces_deleteTableCmd.MarkFlagRequired("table-name")
	keyspacesCmd.AddCommand(keyspaces_deleteTableCmd)
}
