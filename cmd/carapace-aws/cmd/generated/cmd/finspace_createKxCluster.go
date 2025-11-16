package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_createKxClusterCmd = &cobra.Command{
	Use:   "create-kx-cluster",
	Short: "Creates a new kdb cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_createKxClusterCmd).Standalone()

	finspace_createKxClusterCmd.Flags().String("auto-scaling-configuration", "", "The configuration based on which FinSpace will scale in or scale out nodes in your cluster.")
	finspace_createKxClusterCmd.Flags().String("availability-zone-id", "", "The availability zone identifiers for the requested regions.")
	finspace_createKxClusterCmd.Flags().String("az-mode", "", "The number of availability zones you want to assign per cluster.")
	finspace_createKxClusterCmd.Flags().String("cache-storage-configurations", "", "The configurations for a read only cache storage associated with a cluster.")
	finspace_createKxClusterCmd.Flags().String("capacity-configuration", "", "A structure for the metadata of a cluster.")
	finspace_createKxClusterCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspace_createKxClusterCmd.Flags().String("cluster-description", "", "A description of the cluster.")
	finspace_createKxClusterCmd.Flags().String("cluster-name", "", "A unique name for the cluster that you want to create.")
	finspace_createKxClusterCmd.Flags().String("cluster-type", "", "Specifies the type of KDB database that is being created.")
	finspace_createKxClusterCmd.Flags().String("code", "", "The details of the custom code that you want to use inside a cluster when analyzing a data.")
	finspace_createKxClusterCmd.Flags().String("command-line-arguments", "", "Defines the key-value pairs to make them available inside the cluster.")
	finspace_createKxClusterCmd.Flags().String("databases", "", "A list of databases that will be available for querying.")
	finspace_createKxClusterCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment.")
	finspace_createKxClusterCmd.Flags().String("execution-role", "", "An IAM role that defines a set of permissions associated with a cluster.")
	finspace_createKxClusterCmd.Flags().String("initialization-script", "", "Specifies a Q program that will be run at launch of a cluster.")
	finspace_createKxClusterCmd.Flags().String("release-label", "", "The version of FinSpace managed kdb to run.")
	finspace_createKxClusterCmd.Flags().String("savedown-storage-configuration", "", "The size and type of the temporary storage that is used to hold data during the savedown process.")
	finspace_createKxClusterCmd.Flags().String("scaling-group-configuration", "", "The structure that stores the configuration details of a scaling group.")
	finspace_createKxClusterCmd.Flags().String("tags", "", "A list of key-value pairs to label the cluster.")
	finspace_createKxClusterCmd.Flags().String("tickerplant-log-configuration", "", "A configuration to store Tickerplant logs.")
	finspace_createKxClusterCmd.Flags().String("vpc-configuration", "", "Configuration details about the network where the Privatelink endpoint of the cluster resides.")
	finspace_createKxClusterCmd.MarkFlagRequired("az-mode")
	finspace_createKxClusterCmd.MarkFlagRequired("cluster-name")
	finspace_createKxClusterCmd.MarkFlagRequired("cluster-type")
	finspace_createKxClusterCmd.MarkFlagRequired("environment-id")
	finspace_createKxClusterCmd.MarkFlagRequired("release-label")
	finspace_createKxClusterCmd.MarkFlagRequired("vpc-configuration")
	finspaceCmd.AddCommand(finspace_createKxClusterCmd)
}
