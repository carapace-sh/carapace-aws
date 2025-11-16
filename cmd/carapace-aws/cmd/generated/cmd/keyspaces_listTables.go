package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_listTablesCmd = &cobra.Command{
	Use:   "list-tables",
	Short: "The `ListTables` operation returns a list of tables for a specified keyspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_listTablesCmd).Standalone()

	keyspaces_listTablesCmd.Flags().String("keyspace-name", "", "The name of the keyspace.")
	keyspaces_listTablesCmd.Flags().String("max-results", "", "The total number of tables to return in the output.")
	keyspaces_listTablesCmd.Flags().String("next-token", "", "The pagination token.")
	keyspaces_listTablesCmd.MarkFlagRequired("keyspace-name")
	keyspacesCmd.AddCommand(keyspaces_listTablesCmd)
}
