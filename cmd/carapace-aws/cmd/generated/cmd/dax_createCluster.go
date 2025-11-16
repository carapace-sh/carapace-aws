package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dax_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates a DAX cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dax_createClusterCmd).Standalone()

	dax_createClusterCmd.Flags().String("availability-zones", "", "The Availability Zones (AZs) in which the cluster nodes will reside after the cluster has been created or updated.")
	dax_createClusterCmd.Flags().String("cluster-endpoint-encryption-type", "", "The type of encryption the cluster's endpoint should support.")
	dax_createClusterCmd.Flags().String("cluster-name", "", "The cluster identifier.")
	dax_createClusterCmd.Flags().String("description", "", "A description of the cluster.")
	dax_createClusterCmd.Flags().String("iam-role-arn", "", "A valid Amazon Resource Name (ARN) that identifies an IAM role.")
	dax_createClusterCmd.Flags().String("network-type", "", "Specifies the IP protocol(s) the cluster uses for network communications.")
	dax_createClusterCmd.Flags().String("node-type", "", "The compute and memory capacity of the nodes in the cluster.")
	dax_createClusterCmd.Flags().String("notification-topic-arn", "", "The Amazon Resource Name (ARN) of the Amazon SNS topic to which notifications will be sent.")
	dax_createClusterCmd.Flags().String("parameter-group-name", "", "The parameter group to be associated with the DAX cluster.")
	dax_createClusterCmd.Flags().String("preferred-maintenance-window", "", "Specifies the weekly time range during which maintenance on the DAX cluster is performed.")
	dax_createClusterCmd.Flags().String("replication-factor", "", "The number of nodes in the DAX cluster.")
	dax_createClusterCmd.Flags().String("security-group-ids", "", "A list of security group IDs to be assigned to each node in the DAX cluster.")
	dax_createClusterCmd.Flags().String("ssespecification", "", "Represents the settings used to enable server-side encryption on the cluster.")
	dax_createClusterCmd.Flags().String("subnet-group-name", "", "The name of the subnet group to be used for the replication group.")
	dax_createClusterCmd.Flags().String("tags", "", "A set of tags to associate with the DAX cluster.")
	dax_createClusterCmd.MarkFlagRequired("cluster-name")
	dax_createClusterCmd.MarkFlagRequired("iam-role-arn")
	dax_createClusterCmd.MarkFlagRequired("node-type")
	dax_createClusterCmd.MarkFlagRequired("replication-factor")
	daxCmd.AddCommand(dax_createClusterCmd)
}
