package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteNodeCmd = &cobra.Command{
	Use:   "delete-node",
	Short: "Delete a Node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteNodeCmd).Standalone()

	medialive_deleteNodeCmd.Flags().String("cluster-id", "", "The ID of the cluster")
	medialive_deleteNodeCmd.Flags().String("node-id", "", "The ID of the node.")
	medialive_deleteNodeCmd.MarkFlagRequired("cluster-id")
	medialive_deleteNodeCmd.MarkFlagRequired("node-id")
	medialiveCmd.AddCommand(medialive_deleteNodeCmd)
}
