package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateNodeCmd = &cobra.Command{
	Use:   "update-node",
	Short: "Change the settings for a Node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateNodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_updateNodeCmd).Standalone()

		medialive_updateNodeCmd.Flags().String("cluster-id", "", "The ID of the cluster")
		medialive_updateNodeCmd.Flags().String("name", "", "Include this parameter only if you want to change the current name of the Node.")
		medialive_updateNodeCmd.Flags().String("node-id", "", "The ID of the node.")
		medialive_updateNodeCmd.Flags().String("role", "", "The initial role of the Node in the Cluster.")
		medialive_updateNodeCmd.Flags().String("sdi-source-mappings", "", "The mappings of a SDI capture card port to a logical SDI data stream")
		medialive_updateNodeCmd.MarkFlagRequired("cluster-id")
		medialive_updateNodeCmd.MarkFlagRequired("node-id")
	})
	medialiveCmd.AddCommand(medialive_updateNodeCmd)
}
