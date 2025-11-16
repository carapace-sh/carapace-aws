package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateNodeStateCmd = &cobra.Command{
	Use:   "update-node-state",
	Short: "Update the state of a node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateNodeStateCmd).Standalone()

	medialive_updateNodeStateCmd.Flags().String("cluster-id", "", "The ID of the cluster")
	medialive_updateNodeStateCmd.Flags().String("node-id", "", "The ID of the node.")
	medialive_updateNodeStateCmd.Flags().String("state", "", "The state to apply to the Node.")
	medialive_updateNodeStateCmd.MarkFlagRequired("cluster-id")
	medialive_updateNodeStateCmd.MarkFlagRequired("node-id")
	medialiveCmd.AddCommand(medialive_updateNodeStateCmd)
}
