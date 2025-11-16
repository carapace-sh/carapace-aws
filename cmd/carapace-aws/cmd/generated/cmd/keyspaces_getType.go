package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_getTypeCmd = &cobra.Command{
	Use:   "get-type",
	Short: "The `GetType` operation returns information about the type, for example the field definitions, the timestamp when the type was last modified, the level of nesting, the status, and details about if the type is used in other types and tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_getTypeCmd).Standalone()

	keyspaces_getTypeCmd.Flags().String("keyspace-name", "", "The name of the keyspace that contains this type.")
	keyspaces_getTypeCmd.Flags().String("type-name", "", "The formatted name of the type.")
	keyspaces_getTypeCmd.MarkFlagRequired("keyspace-name")
	keyspaces_getTypeCmd.MarkFlagRequired("type-name")
	keyspacesCmd.AddCommand(keyspaces_getTypeCmd)
}
