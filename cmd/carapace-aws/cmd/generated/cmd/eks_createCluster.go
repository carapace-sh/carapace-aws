package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var eks_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates an Amazon EKS control plane.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(eks_createClusterCmd).Standalone()

	eks_createClusterCmd.Flags().String("access-config", "", "The access configuration for the cluster.")
	eks_createClusterCmd.Flags().String("bootstrap-self-managed-addons", "", "If you set this value to `False` when creating a cluster, the default networking add-ons will not be installed.")
	eks_createClusterCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
	eks_createClusterCmd.Flags().String("compute-config", "", "Enable or disable the compute capability of EKS Auto Mode when creating your EKS Auto Mode cluster.")
	eks_createClusterCmd.Flags().String("deletion-protection", "", "Indicates whether to enable deletion protection for the cluster.")
	eks_createClusterCmd.Flags().String("encryption-config", "", "The encryption configuration for the cluster.")
	eks_createClusterCmd.Flags().String("kubernetes-network-config", "", "The Kubernetes network configuration for the cluster.")
	eks_createClusterCmd.Flags().String("logging", "", "Enable or disable exporting the Kubernetes control plane logs for your cluster to CloudWatch Logs .")
	eks_createClusterCmd.Flags().String("name", "", "The unique name to give to your cluster.")
	eks_createClusterCmd.Flags().String("outpost-config", "", "An object representing the configuration of your local Amazon EKS cluster on an Amazon Web Services Outpost.")
	eks_createClusterCmd.Flags().String("remote-network-config", "", "The configuration in the cluster for EKS Hybrid Nodes.")
	eks_createClusterCmd.Flags().String("resources-vpc-config", "", "The VPC configuration that's used by the cluster control plane.")
	eks_createClusterCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role that provides permissions for the Kubernetes control plane to make calls to Amazon Web Services API operations on your behalf.")
	eks_createClusterCmd.Flags().String("storage-config", "", "Enable or disable the block storage capability of EKS Auto Mode when creating your EKS Auto Mode cluster.")
	eks_createClusterCmd.Flags().String("tags", "", "Metadata that assists with categorization and organization.")
	eks_createClusterCmd.Flags().String("upgrade-policy", "", "New clusters, by default, have extended support enabled.")
	eks_createClusterCmd.Flags().String("version", "", "The desired Kubernetes version for your cluster.")
	eks_createClusterCmd.Flags().String("zonal-shift-config", "", "Enable or disable ARC zonal shift for the cluster.")
	eks_createClusterCmd.MarkFlagRequired("name")
	eks_createClusterCmd.MarkFlagRequired("resources-vpc-config")
	eks_createClusterCmd.MarkFlagRequired("role-arn")
	eksCmd.AddCommand(eks_createClusterCmd)
}
