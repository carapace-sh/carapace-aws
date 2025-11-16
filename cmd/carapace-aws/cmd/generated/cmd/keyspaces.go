package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspacesCmd = &cobra.Command{
	Use:   "keyspaces",
	Short: "Amazon Keyspaces (for Apache Cassandra) is a scalable, highly available, and managed Apache Cassandra-compatible database service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspacesCmd).Standalone()

	rootCmd.AddCommand(keyspacesCmd)
}
