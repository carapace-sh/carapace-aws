package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_createKeyspaceCmd = &cobra.Command{
	Use:   "create-keyspace",
	Short: "The `CreateKeyspace` operation adds a new keyspace to your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_createKeyspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspaces_createKeyspaceCmd).Standalone()

		keyspaces_createKeyspaceCmd.Flags().String("keyspace-name", "", "The name of the keyspace to be created.")
		keyspaces_createKeyspaceCmd.Flags().String("replication-specification", "", "The replication specification of the keyspace includes:")
		keyspaces_createKeyspaceCmd.Flags().String("tags", "", "A list of key-value pair tags to be attached to the keyspace.")
		keyspaces_createKeyspaceCmd.MarkFlagRequired("keyspace-name")
	})
	keyspacesCmd.AddCommand(keyspaces_createKeyspaceCmd)
}
