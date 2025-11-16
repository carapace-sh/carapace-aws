package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_updateKeyspaceCmd = &cobra.Command{
	Use:   "update-keyspace",
	Short: "Adds a new Amazon Web Services Region to the keyspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_updateKeyspaceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspaces_updateKeyspaceCmd).Standalone()

		keyspaces_updateKeyspaceCmd.Flags().String("client-side-timestamps", "", "")
		keyspaces_updateKeyspaceCmd.Flags().String("keyspace-name", "", "The name of the keyspace.")
		keyspaces_updateKeyspaceCmd.Flags().String("replication-specification", "", "")
		keyspaces_updateKeyspaceCmd.MarkFlagRequired("keyspace-name")
		keyspaces_updateKeyspaceCmd.MarkFlagRequired("replication-specification")
	})
	keyspacesCmd.AddCommand(keyspaces_updateKeyspaceCmd)
}
