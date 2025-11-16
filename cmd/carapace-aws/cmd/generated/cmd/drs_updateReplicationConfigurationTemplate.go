package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_updateReplicationConfigurationTemplateCmd = &cobra.Command{
	Use:   "update-replication-configuration-template",
	Short: "Updates a ReplicationConfigurationTemplate by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_updateReplicationConfigurationTemplateCmd).Standalone()

	drs_updateReplicationConfigurationTemplateCmd.Flags().String("arn", "", "The Replication Configuration Template ARN.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().Bool("associate-default-security-group", false, "Whether to associate the default Elastic Disaster Recovery Security group with the Replication Configuration Template.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().Bool("auto-replicate-new-disks", false, "Whether to allow the AWS replication agent to automatically replicate newly added disks.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().String("bandwidth-throttling", "", "Configure bandwidth throttling for the outbound data transfer rate of the Source Server in Mbps.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().Bool("create-public-ip", false, "Whether to create a Public IP for the Recovery Instance by default.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().String("data-plane-routing", "", "The data plane routing mechanism that will be used for replication.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().String("default-large-staging-disk-type", "", "The Staging Disk EBS volume type to be used during replication.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().String("ebs-encryption", "", "The type of EBS encryption to be used during replication.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().String("ebs-encryption-key-arn", "", "The ARN of the EBS encryption key to be used during replication.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().Bool("no-associate-default-security-group", false, "Whether to associate the default Elastic Disaster Recovery Security group with the Replication Configuration Template.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().Bool("no-auto-replicate-new-disks", false, "Whether to allow the AWS replication agent to automatically replicate newly added disks.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().Bool("no-create-public-ip", false, "Whether to create a Public IP for the Recovery Instance by default.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().Bool("no-use-dedicated-replication-server", false, "Whether to use a dedicated Replication Server in the replication staging area.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().String("pit-policy", "", "The Point in time (PIT) policy to manage snapshots taken during replication.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().String("replication-configuration-template-id", "", "The Replication Configuration Template ID.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().String("replication-server-instance-type", "", "The instance type to be used for the replication server.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().String("replication-servers-security-groups-ids", "", "The security group IDs that will be used by the replication server.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().String("staging-area-subnet-id", "", "The subnet to be used by the replication staging area.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().String("staging-area-tags", "", "A set of tags to be associated with all resources created in the replication staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.")
	drs_updateReplicationConfigurationTemplateCmd.Flags().Bool("use-dedicated-replication-server", false, "Whether to use a dedicated Replication Server in the replication staging area.")
	drs_updateReplicationConfigurationTemplateCmd.Flag("no-associate-default-security-group").Hidden = true
	drs_updateReplicationConfigurationTemplateCmd.Flag("no-auto-replicate-new-disks").Hidden = true
	drs_updateReplicationConfigurationTemplateCmd.Flag("no-create-public-ip").Hidden = true
	drs_updateReplicationConfigurationTemplateCmd.Flag("no-use-dedicated-replication-server").Hidden = true
	drs_updateReplicationConfigurationTemplateCmd.MarkFlagRequired("replication-configuration-template-id")
	drsCmd.AddCommand(drs_updateReplicationConfigurationTemplateCmd)
}
