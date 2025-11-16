package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspaces_getKeyspaceCmd = &cobra.Command{
	Use:   "get-keyspace",
	Short: "Returns the name of the specified keyspace, the Amazon Resource Name (ARN), the replication strategy, the Amazon Web Services Regions of a multi-Region keyspace, and the status of newly added Regions after an `UpdateKeyspace` operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspaces_getKeyspaceCmd).Standalone()

	keyspaces_getKeyspaceCmd.Flags().String("keyspace-name", "", "The name of the keyspace.")
	keyspaces_getKeyspaceCmd.MarkFlagRequired("keyspace-name")
	keyspacesCmd.AddCommand(keyspaces_getKeyspaceCmd)
}
