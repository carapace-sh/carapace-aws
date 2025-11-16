package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyClusterCmd = &cobra.Command{
	Use:   "modify-cluster",
	Short: "Modifies the settings for a cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyClusterCmd).Standalone()

	redshift_modifyClusterCmd.Flags().String("allow-version-upgrade", "", "If `true`, major version upgrades will be applied automatically to the cluster during the maintenance window.")
	redshift_modifyClusterCmd.Flags().String("automated-snapshot-retention-period", "", "The number of days that automated snapshots are retained.")
	redshift_modifyClusterCmd.Flags().String("availability-zone", "", "The option to initiate relocation for an Amazon Redshift cluster to the target Availability Zone.")
	redshift_modifyClusterCmd.Flags().String("availability-zone-relocation", "", "The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster modification is complete.")
	redshift_modifyClusterCmd.Flags().String("cluster-identifier", "", "The unique identifier of the cluster to be modified.")
	redshift_modifyClusterCmd.Flags().String("cluster-parameter-group-name", "", "The name of the cluster parameter group to apply to this cluster.")
	redshift_modifyClusterCmd.Flags().String("cluster-security-groups", "", "A list of cluster security groups to be authorized on this cluster.")
	redshift_modifyClusterCmd.Flags().String("cluster-type", "", "The new cluster type.")
	redshift_modifyClusterCmd.Flags().String("cluster-version", "", "The new version number of the Amazon Redshift engine to upgrade to.")
	redshift_modifyClusterCmd.Flags().String("elastic-ip", "", "The Elastic IP (EIP) address for the cluster.")
	redshift_modifyClusterCmd.Flags().String("encrypted", "", "Indicates whether the cluster is encrypted.")
	redshift_modifyClusterCmd.Flags().String("enhanced-vpc-routing", "", "An option that specifies whether to create the cluster with enhanced VPC routing enabled.")
	redshift_modifyClusterCmd.Flags().String("hsm-client-certificate-identifier", "", "Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM.")
	redshift_modifyClusterCmd.Flags().String("hsm-configuration-identifier", "", "Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.")
	redshift_modifyClusterCmd.Flags().String("ip-address-type", "", "The IP address types that the cluster supports.")
	redshift_modifyClusterCmd.Flags().String("kms-key-id", "", "The Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.")
	redshift_modifyClusterCmd.Flags().String("maintenance-track-name", "", "The name for the maintenance track that you want to assign for the cluster.")
	redshift_modifyClusterCmd.Flags().String("manage-master-password", "", "If `true`, Amazon Redshift uses Secrets Manager to manage this cluster's admin credentials.")
	redshift_modifyClusterCmd.Flags().String("manual-snapshot-retention-period", "", "The default for number of days that a newly created manual snapshot is retained.")
	redshift_modifyClusterCmd.Flags().String("master-password-secret-kms-key-id", "", "The ID of the Key Management Service (KMS) key used to encrypt and store the cluster's admin credentials secret.")
	redshift_modifyClusterCmd.Flags().String("master-user-password", "", "The new password for the cluster admin user.")
	redshift_modifyClusterCmd.Flags().String("multi-az", "", "If true and the cluster is currently only deployed in a single Availability Zone, the cluster will be modified to be deployed in two Availability Zones.")
	redshift_modifyClusterCmd.Flags().String("new-cluster-identifier", "", "The new identifier for the cluster.")
	redshift_modifyClusterCmd.Flags().String("node-type", "", "The new node type of the cluster.")
	redshift_modifyClusterCmd.Flags().String("number-of-nodes", "", "The new number of nodes of the cluster.")
	redshift_modifyClusterCmd.Flags().String("port", "", "The option to change the port of an Amazon Redshift cluster.")
	redshift_modifyClusterCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range (in UTC) during which system maintenance can occur, if necessary.")
	redshift_modifyClusterCmd.Flags().String("publicly-accessible", "", "If `true`, the cluster can be accessed from a public network.")
	redshift_modifyClusterCmd.Flags().String("vpc-security-group-ids", "", "A list of virtual private cloud (VPC) security groups to be associated with the cluster.")
	redshift_modifyClusterCmd.MarkFlagRequired("cluster-identifier")
	redshiftCmd.AddCommand(redshift_modifyClusterCmd)
}
