package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createClusterCmd = &cobra.Command{
	Use:   "create-cluster",
	Short: "Creates a new cluster with the specified parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_createClusterCmd).Standalone()

		redshift_createClusterCmd.Flags().String("additional-info", "", "Reserved.")
		redshift_createClusterCmd.Flags().String("allow-version-upgrade", "", "If `true`, major version upgrades can be applied during the maintenance window to the Amazon Redshift engine that is running on the cluster.")
		redshift_createClusterCmd.Flags().String("aqua-configuration-status", "", "This parameter is retired.")
		redshift_createClusterCmd.Flags().String("automated-snapshot-retention-period", "", "The number of days that automated snapshots are retained.")
		redshift_createClusterCmd.Flags().String("availability-zone", "", "The EC2 Availability Zone (AZ) in which you want Amazon Redshift to provision the cluster.")
		redshift_createClusterCmd.Flags().String("availability-zone-relocation", "", "The option to enable relocation for an Amazon Redshift cluster between Availability Zones after the cluster is created.")
		redshift_createClusterCmd.Flags().String("cluster-identifier", "", "A unique identifier for the cluster.")
		redshift_createClusterCmd.Flags().String("cluster-parameter-group-name", "", "The name of the parameter group to be associated with this cluster.")
		redshift_createClusterCmd.Flags().String("cluster-security-groups", "", "A list of security groups to be associated with this cluster.")
		redshift_createClusterCmd.Flags().String("cluster-subnet-group-name", "", "The name of a cluster subnet group to be associated with this cluster.")
		redshift_createClusterCmd.Flags().String("cluster-type", "", "The type of the cluster.")
		redshift_createClusterCmd.Flags().String("cluster-version", "", "The version of the Amazon Redshift engine software that you want to deploy on the cluster.")
		redshift_createClusterCmd.Flags().String("dbname", "", "The name of the first database to be created when the cluster is created.")
		redshift_createClusterCmd.Flags().String("default-iam-role-arn", "", "The Amazon Resource Name (ARN) for the IAM role that was set as default for the cluster when the cluster was created.")
		redshift_createClusterCmd.Flags().String("elastic-ip", "", "The Elastic IP (EIP) address for the cluster.")
		redshift_createClusterCmd.Flags().String("encrypted", "", "If `true`, the data in the cluster is encrypted at rest.")
		redshift_createClusterCmd.Flags().String("enhanced-vpc-routing", "", "An option that specifies whether to create the cluster with enhanced VPC routing enabled.")
		redshift_createClusterCmd.Flags().String("hsm-client-certificate-identifier", "", "Specifies the name of the HSM client certificate the Amazon Redshift cluster uses to retrieve the data encryption keys stored in an HSM.")
		redshift_createClusterCmd.Flags().String("hsm-configuration-identifier", "", "Specifies the name of the HSM configuration that contains the information the Amazon Redshift cluster can use to retrieve and store keys in an HSM.")
		redshift_createClusterCmd.Flags().String("iam-roles", "", "A list of Identity and Access Management (IAM) roles that can be used by the cluster to access other Amazon Web Services services.")
		redshift_createClusterCmd.Flags().String("ip-address-type", "", "The IP address types that the cluster supports.")
		redshift_createClusterCmd.Flags().String("kms-key-id", "", "The Key Management Service (KMS) key ID of the encryption key that you want to use to encrypt data in the cluster.")
		redshift_createClusterCmd.Flags().String("load-sample-data", "", "A flag that specifies whether to load sample data once the cluster is created.")
		redshift_createClusterCmd.Flags().String("maintenance-track-name", "", "An optional parameter for the name of the maintenance track for the cluster.")
		redshift_createClusterCmd.Flags().String("manage-master-password", "", "If `true`, Amazon Redshift uses Secrets Manager to manage this cluster's admin credentials.")
		redshift_createClusterCmd.Flags().String("manual-snapshot-retention-period", "", "The default number of days to retain a manual snapshot.")
		redshift_createClusterCmd.Flags().String("master-password-secret-kms-key-id", "", "The ID of the Key Management Service (KMS) key used to encrypt and store the cluster's admin credentials secret.")
		redshift_createClusterCmd.Flags().String("master-user-password", "", "The password associated with the admin user account for the cluster that is being created.")
		redshift_createClusterCmd.Flags().String("master-username", "", "The user name associated with the admin user account for the cluster that is being created.")
		redshift_createClusterCmd.Flags().String("multi-az", "", "If true, Amazon Redshift will deploy the cluster in two Availability Zones (AZ).")
		redshift_createClusterCmd.Flags().String("node-type", "", "The node type to be provisioned for the cluster.")
		redshift_createClusterCmd.Flags().String("number-of-nodes", "", "The number of compute nodes in the cluster.")
		redshift_createClusterCmd.Flags().String("port", "", "The port number on which the cluster accepts incoming connections.")
		redshift_createClusterCmd.Flags().String("preferred-maintenance-window", "", "The weekly time range (in UTC) during which automated cluster maintenance can occur.")
		redshift_createClusterCmd.Flags().String("publicly-accessible", "", "If `true`, the cluster can be accessed from a public network.")
		redshift_createClusterCmd.Flags().String("redshift-idc-application-arn", "", "The Amazon resource name (ARN) of the Amazon Redshift IAM Identity Center application.")
		redshift_createClusterCmd.Flags().String("snapshot-schedule-identifier", "", "A unique identifier for the snapshot schedule.")
		redshift_createClusterCmd.Flags().String("tags", "", "A list of tag instances.")
		redshift_createClusterCmd.Flags().String("vpc-security-group-ids", "", "A list of Virtual Private Cloud (VPC) security groups to be associated with the cluster.")
		redshift_createClusterCmd.MarkFlagRequired("cluster-identifier")
		redshift_createClusterCmd.MarkFlagRequired("master-username")
		redshift_createClusterCmd.MarkFlagRequired("node-type")
	})
	redshiftCmd.AddCommand(redshift_createClusterCmd)
}
