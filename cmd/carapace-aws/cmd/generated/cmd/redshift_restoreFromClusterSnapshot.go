package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_restoreFromClusterSnapshotCmd = &cobra.Command{
	Use:   "restore-from-cluster-snapshot",
	Short: "Creates a new cluster from a snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_restoreFromClusterSnapshotCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_restoreFromClusterSnapshotCmd).Standalone()

		redshift_restoreFromClusterSnapshotCmd.Flags().String("additional-info", "", "Reserved.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("allow-version-upgrade", "", "If `true`, major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("aqua-configuration-status", "", "This parameter is retired.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("automated-snapshot-retention-period", "", "The number of days that automated snapshots are retained.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("availability-zone", "", "The Amazon EC2 Availability Zone in which to restore the cluster.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("availability-zone-relocation", "", "The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster is restored.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("cluster-identifier", "", "The identifier of the cluster that will be created from restoring the snapshot.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("cluster-parameter-group-name", "", "The name of the parameter group to be associated with this cluster.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("cluster-security-groups", "", "A list of security groups to be associated with this cluster.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("cluster-subnet-group-name", "", "The name of the subnet group where you want to cluster restored.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("default-iam-role-arn", "", "The Amazon Resource Name (ARN) for the IAM role that was set as default for the cluster when the cluster was last modified while it was restored from a snapshot.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("elastic-ip", "", "The Elastic IP (EIP) address for the cluster.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("encrypted", "", "Enables support for restoring an unencrypted snapshot to a cluster encrypted with Key Management Service (KMS) and a customer managed key.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("enhanced-vpc-routing", "", "An option that specifies whether to create the cluster with enhanced VPC routing enabled.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("hsm-client-certificate-identifier", "", "Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("hsm-configuration-identifier", "", "Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("iam-roles", "", "A list of Identity and Access Management (IAM) roles that can be used by the cluster to access other Amazon Web Services services.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("ip-address-type", "", "The IP address type for the cluster.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("kms-key-id", "", "The Key Management Service (KMS) key ID of the encryption key that encrypts data in the cluster restored from a shared snapshot.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("maintenance-track-name", "", "The name of the maintenance track for the restored cluster.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("manage-master-password", "", "If `true`, Amazon Redshift uses Secrets Manager to manage the restored cluster's admin credentials.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("manual-snapshot-retention-period", "", "The default number of days to retain a manual snapshot.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("master-password-secret-kms-key-id", "", "The ID of the Key Management Service (KMS) key used to encrypt and store the cluster's admin credentials secret.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("multi-az", "", "If true, the snapshot will be restored to a cluster deployed in two Availability Zones.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("node-type", "", "The node type that the restored cluster will be provisioned with.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("number-of-nodes", "", "The number of nodes specified when provisioning the restored cluster.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("owner-account", "", "The Amazon Web Services account used to create or copy the snapshot.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("port", "", "The port number on which the cluster accepts connections.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range (in UTC) during which automated cluster maintenance can occur.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("publicly-accessible", "", "If `true`, the cluster can be accessed from a public network.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("reserved-node-id", "", "The identifier of the target reserved node offering.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("snapshot-arn", "", "The Amazon Resource Name (ARN) of the snapshot associated with the message to restore from a cluster.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("snapshot-cluster-identifier", "", "The name of the cluster the source snapshot was created from.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("snapshot-identifier", "", "The name of the snapshot from which to create the new cluster.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("snapshot-schedule-identifier", "", "A unique identifier for the snapshot schedule.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("target-reserved-node-offering-id", "", "The identifier of the target reserved node offering.")
		redshift_restoreFromClusterSnapshotCmd.Flags().String("vpc-security-group-ids", "", "A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.")
		redshift_restoreFromClusterSnapshotCmd.MarkFlagRequired("cluster-identifier")
	})
	redshiftCmd.AddCommand(redshift_restoreFromClusterSnapshotCmd)
}
