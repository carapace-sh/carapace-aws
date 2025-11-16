package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_createReplicationConfigurationTemplateCmd = &cobra.Command{
	Use:   "create-replication-configuration-template",
	Short: "Creates a new ReplicationConfigurationTemplate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_createReplicationConfigurationTemplateCmd).Standalone()

	drs_createReplicationConfigurationTemplateCmd.Flags().Bool("associate-default-security-group", false, "Whether to associate the default Elastic Disaster Recovery Security group with the Replication Configuration Template.")
	drs_createReplicationConfigurationTemplateCmd.Flags().Bool("auto-replicate-new-disks", false, "Whether to allow the AWS replication agent to automatically replicate newly added disks.")
	drs_createReplicationConfigurationTemplateCmd.Flags().String("bandwidth-throttling", "", "Configure bandwidth throttling for the outbound data transfer rate of the Source Server in Mbps.")
	drs_createReplicationConfigurationTemplateCmd.Flags().Bool("create-public-ip", false, "Whether to create a Public IP for the Recovery Instance by default.")
	drs_createReplicationConfigurationTemplateCmd.Flags().String("data-plane-routing", "", "The data plane routing mechanism that will be used for replication.")
	drs_createReplicationConfigurationTemplateCmd.Flags().String("default-large-staging-disk-type", "", "The Staging Disk EBS volume type to be used during replication.")
	drs_createReplicationConfigurationTemplateCmd.Flags().String("ebs-encryption", "", "The type of EBS encryption to be used during replication.")
	drs_createReplicationConfigurationTemplateCmd.Flags().String("ebs-encryption-key-arn", "", "The ARN of the EBS encryption key to be used during replication.")
	drs_createReplicationConfigurationTemplateCmd.Flags().Bool("no-associate-default-security-group", false, "Whether to associate the default Elastic Disaster Recovery Security group with the Replication Configuration Template.")
	drs_createReplicationConfigurationTemplateCmd.Flags().Bool("no-auto-replicate-new-disks", false, "Whether to allow the AWS replication agent to automatically replicate newly added disks.")
	drs_createReplicationConfigurationTemplateCmd.Flags().Bool("no-create-public-ip", false, "Whether to create a Public IP for the Recovery Instance by default.")
	drs_createReplicationConfigurationTemplateCmd.Flags().Bool("no-use-dedicated-replication-server", false, "Whether to use a dedicated Replication Server in the replication staging area.")
	drs_createReplicationConfigurationTemplateCmd.Flags().String("pit-policy", "", "The Point in time (PIT) policy to manage snapshots taken during replication.")
	drs_createReplicationConfigurationTemplateCmd.Flags().String("replication-server-instance-type", "", "The instance type to be used for the replication server.")
	drs_createReplicationConfigurationTemplateCmd.Flags().String("replication-servers-security-groups-ids", "", "The security group IDs that will be used by the replication server.")
	drs_createReplicationConfigurationTemplateCmd.Flags().String("staging-area-subnet-id", "", "The subnet to be used by the replication staging area.")
	drs_createReplicationConfigurationTemplateCmd.Flags().String("staging-area-tags", "", "A set of tags to be associated with all resources created in the replication staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.")
	drs_createReplicationConfigurationTemplateCmd.Flags().String("tags", "", "A set of tags to be associated with the Replication Configuration Template resource.")
	drs_createReplicationConfigurationTemplateCmd.Flags().Bool("use-dedicated-replication-server", false, "Whether to use a dedicated Replication Server in the replication staging area.")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("associate-default-security-group")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("bandwidth-throttling")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("create-public-ip")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("data-plane-routing")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("default-large-staging-disk-type")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("ebs-encryption")
	drs_createReplicationConfigurationTemplateCmd.Flag("no-associate-default-security-group").Hidden = true
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("no-associate-default-security-group")
	drs_createReplicationConfigurationTemplateCmd.Flag("no-auto-replicate-new-disks").Hidden = true
	drs_createReplicationConfigurationTemplateCmd.Flag("no-create-public-ip").Hidden = true
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("no-create-public-ip")
	drs_createReplicationConfigurationTemplateCmd.Flag("no-use-dedicated-replication-server").Hidden = true
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("no-use-dedicated-replication-server")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("pit-policy")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("replication-server-instance-type")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("replication-servers-security-groups-ids")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("staging-area-subnet-id")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("staging-area-tags")
	drs_createReplicationConfigurationTemplateCmd.MarkFlagRequired("use-dedicated-replication-server")
	drsCmd.AddCommand(drs_createReplicationConfigurationTemplateCmd)
}
