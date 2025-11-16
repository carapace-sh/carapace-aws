package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateClusterCmd = &cobra.Command{
	Use:   "update-cluster",
	Short: "Change the settings for a Cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateClusterCmd).Standalone()

	medialive_updateClusterCmd.Flags().String("cluster-id", "", "The ID of the cluster")
	medialive_updateClusterCmd.Flags().String("name", "", "Include this parameter only if you want to change the current name of the Cluster.")
	medialive_updateClusterCmd.Flags().String("network-settings", "", "Include this property only if you want to change the current connections between the Nodes in the Cluster and the Networks the Cluster is associated with.")
	medialive_updateClusterCmd.MarkFlagRequired("cluster-id")
	medialiveCmd.AddCommand(medialive_updateClusterCmd)
}
