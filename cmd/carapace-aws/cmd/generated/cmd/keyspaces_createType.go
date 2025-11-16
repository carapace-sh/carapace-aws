package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_createTypeCmd = &cobra.Command{
	Use:   "create-type",
	Short: "The `CreateType` operation creates a new user-defined type in the specified keyspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_createTypeCmd).Standalone()

	keyspaces_createTypeCmd.Flags().String("field-definitions", "", "The field definitions, consisting of names and types, that define this type.")
	keyspaces_createTypeCmd.Flags().String("keyspace-name", "", "The name of the keyspace.")
	keyspaces_createTypeCmd.Flags().String("type-name", "", "The name of the user-defined type.")
	keyspaces_createTypeCmd.MarkFlagRequired("field-definitions")
	keyspaces_createTypeCmd.MarkFlagRequired("keyspace-name")
	keyspaces_createTypeCmd.MarkFlagRequired("type-name")
	keyspacesCmd.AddCommand(keyspaces_createTypeCmd)
}
