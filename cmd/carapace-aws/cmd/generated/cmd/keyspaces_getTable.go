package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_getTableCmd = &cobra.Command{
	Use:   "get-table",
	Short: "Returns information about the table, including the table's name and current status, the keyspace name, configuration settings, and metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_getTableCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspaces_getTableCmd).Standalone()

		keyspaces_getTableCmd.Flags().String("keyspace-name", "", "The name of the keyspace that the table is stored in.")
		keyspaces_getTableCmd.Flags().String("table-name", "", "The name of the table.")
		keyspaces_getTableCmd.MarkFlagRequired("keyspace-name")
		keyspaces_getTableCmd.MarkFlagRequired("table-name")
	})
	keyspacesCmd.AddCommand(keyspaces_getTableCmd)
}
