package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_listTypesCmd = &cobra.Command{
	Use:   "list-types",
	Short: "The `ListTypes` operation returns a list of types for a specified keyspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_listTypesCmd).Standalone()

	keyspaces_listTypesCmd.Flags().String("keyspace-name", "", "The name of the keyspace that contains the listed types.")
	keyspaces_listTypesCmd.Flags().String("max-results", "", "The total number of types to return in the output.")
	keyspaces_listTypesCmd.Flags().String("next-token", "", "The pagination token.")
	keyspaces_listTypesCmd.MarkFlagRequired("keyspace-name")
	keyspacesCmd.AddCommand(keyspaces_listTypesCmd)
}
