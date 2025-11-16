package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keyspacesstreamsCmd = &cobra.Command{
	Use:   "keyspacesstreams",
	Short: "Amazon Keyspaces (for Apache Cassandra) change data capture (CDC) records change events for Amazon Keyspaces tables.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyspacesstreamsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(keyspacesstreamsCmd).Standalone()

	})
	rootCmd.AddCommand(keyspacesstreamsCmd)
}
