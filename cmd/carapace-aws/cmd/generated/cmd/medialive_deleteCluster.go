package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Delete a Cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_deleteClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_deleteClusterCmd).Standalone()

		medialive_deleteClusterCmd.Flags().String("cluster-id", "", "The ID of the cluster.")
		medialive_deleteClusterCmd.MarkFlagRequired("cluster-id")
	})
	medialiveCmd.AddCommand(medialive_deleteClusterCmd)
}
