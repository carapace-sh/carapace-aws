package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_deleteNamespaceCmd = &cobra.Command{
	Use:   "delete-namespace",
	Short: "Deletes the specified namespace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_deleteNamespaceCmd).Standalone()

	iotthingsgraphCmd.AddCommand(iotthingsgraph_deleteNamespaceCmd)
}
