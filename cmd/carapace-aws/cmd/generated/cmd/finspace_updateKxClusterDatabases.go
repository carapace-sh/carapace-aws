package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_updateKxClusterDatabasesCmd = &cobra.Command{
	Use:   "update-kx-cluster-databases",
	Short: "Updates the databases mounted on a kdb cluster, which includes the `changesetId` and all the dbPaths to be cached.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_updateKxClusterDatabasesCmd).Standalone()

	finspace_updateKxClusterDatabasesCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_updateKxClusterDatabasesCmd.Flags().String("cluster-name", "", "A unique name for the cluster that you want to modify.")
	finspace_updateKxClusterDatabasesCmd.Flags().String("databases", "", "The structure of databases mounted on the cluster.")
	finspace_updateKxClusterDatabasesCmd.Flags().String("deployment-configuration", "", "The configuration that allows you to choose how you want to update the databases on a cluster.")
	finspace_updateKxClusterDatabasesCmd.Flags().String("environment-id", "", "The unique identifier of a kdb environment.")
	finspace_updateKxClusterDatabasesCmd.MarkFlagRequired("cluster-name")
	finspace_updateKxClusterDatabasesCmd.MarkFlagRequired("databases")
	finspace_updateKxClusterDatabasesCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_updateKxClusterDatabasesCmd)
}
