package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_getNamespaceDeletionStatusCmd = &cobra.Command{
	Use:   "get-namespace-deletion-status",
	Short: "Gets the status of a namespace deletion task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_getNamespaceDeletionStatusCmd).Standalone()

	iotthingsgraphCmd.AddCommand(iotthingsgraph_getNamespaceDeletionStatusCmd)
}
