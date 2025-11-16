package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptune_createGlobalClusterCmd = &cobra.Command{
	Use:   "create-global-cluster",
	Short: "Creates a Neptune global database spread across multiple Amazon Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptune_createGlobalClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptune_createGlobalClusterCmd).Standalone()

		neptune_createGlobalClusterCmd.Flags().String("deletion-protection", "", "The deletion protection setting for the new global database.")
		neptune_createGlobalClusterCmd.Flags().String("engine", "", "The name of the database engine to be used in the global database.")
		neptune_createGlobalClusterCmd.Flags().String("engine-version", "", "The Neptune engine version to be used by the global database.")
		neptune_createGlobalClusterCmd.Flags().String("global-cluster-identifier", "", "The cluster identifier of the new global database cluster.")
		neptune_createGlobalClusterCmd.Flags().String("source-dbcluster-identifier", "", "(*Optional*) The Amazon Resource Name (ARN) of an existing Neptune DB cluster to use as the primary cluster of the new global database.")
		neptune_createGlobalClusterCmd.Flags().String("storage-encrypted", "", "The storage encryption setting for the new global database cluster.")
		neptune_createGlobalClusterCmd.MarkFlagRequired("global-cluster-identifier")
	})
	neptuneCmd.AddCommand(neptune_createGlobalClusterCmd)
}
