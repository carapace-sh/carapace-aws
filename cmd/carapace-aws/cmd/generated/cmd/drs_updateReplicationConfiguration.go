package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_updateReplicationConfigurationCmd = &cobra.Command{
	Use:   "update-replication-configuration",
	Short: "Allows you to update a ReplicationConfiguration by Source Server ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_updateReplicationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(drs_updateReplicationConfigurationCmd).Standalone()

		drs_updateReplicationConfigurationCmd.Flags().Bool("associate-default-security-group", false, "Whether to associate the default Elastic Disaster Recovery Security group with the Replication Configuration.")
		drs_updateReplicationConfigurationCmd.Flags().Bool("auto-replicate-new-disks", false, "Whether to allow the AWS replication agent to automatically replicate newly added disks.")
		drs_updateReplicationConfigurationCmd.Flags().String("bandwidth-throttling", "", "Configure bandwidth throttling for the outbound data transfer rate of the Source Server in Mbps.")
		drs_updateReplicationConfigurationCmd.Flags().Bool("create-public-ip", false, "Whether to create a Public IP for the Recovery Instance by default.")
		drs_updateReplicationConfigurationCmd.Flags().String("data-plane-routing", "", "The data plane routing mechanism that will be used for replication.")
		drs_updateReplicationConfigurationCmd.Flags().String("default-large-staging-disk-type", "", "The Staging Disk EBS volume type to be used during replication.")
		drs_updateReplicationConfigurationCmd.Flags().String("ebs-encryption", "", "The type of EBS encryption to be used during replication.")
		drs_updateReplicationConfigurationCmd.Flags().String("ebs-encryption-key-arn", "", "The ARN of the EBS encryption key to be used during replication.")
		drs_updateReplicationConfigurationCmd.Flags().String("name", "", "The name of the Replication Configuration.")
		drs_updateReplicationConfigurationCmd.Flags().Bool("no-associate-default-security-group", false, "Whether to associate the default Elastic Disaster Recovery Security group with the Replication Configuration.")
		drs_updateReplicationConfigurationCmd.Flags().Bool("no-auto-replicate-new-disks", false, "Whether to allow the AWS replication agent to automatically replicate newly added disks.")
		drs_updateReplicationConfigurationCmd.Flags().Bool("no-create-public-ip", false, "Whether to create a Public IP for the Recovery Instance by default.")
		drs_updateReplicationConfigurationCmd.Flags().Bool("no-use-dedicated-replication-server", false, "Whether to use a dedicated Replication Server in the replication staging area.")
		drs_updateReplicationConfigurationCmd.Flags().String("pit-policy", "", "The Point in time (PIT) policy to manage snapshots taken during replication.")
		drs_updateReplicationConfigurationCmd.Flags().String("replicated-disks", "", "The configuration of the disks of the Source Server to be replicated.")
		drs_updateReplicationConfigurationCmd.Flags().String("replication-server-instance-type", "", "The instance type to be used for the replication server.")
		drs_updateReplicationConfigurationCmd.Flags().String("replication-servers-security-groups-ids", "", "The security group IDs that will be used by the replication server.")
		drs_updateReplicationConfigurationCmd.Flags().String("source-server-id", "", "The ID of the Source Server for this Replication Configuration.")
		drs_updateReplicationConfigurationCmd.Flags().String("staging-area-subnet-id", "", "The subnet to be used by the replication staging area.")
		drs_updateReplicationConfigurationCmd.Flags().String("staging-area-tags", "", "A set of tags to be associated with all resources created in the replication staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.")
		drs_updateReplicationConfigurationCmd.Flags().Bool("use-dedicated-replication-server", false, "Whether to use a dedicated Replication Server in the replication staging area.")
		drs_updateReplicationConfigurationCmd.Flag("no-associate-default-security-group").Hidden = true
		drs_updateReplicationConfigurationCmd.Flag("no-auto-replicate-new-disks").Hidden = true
		drs_updateReplicationConfigurationCmd.Flag("no-create-public-ip").Hidden = true
		drs_updateReplicationConfigurationCmd.Flag("no-use-dedicated-replication-server").Hidden = true
		drs_updateReplicationConfigurationCmd.MarkFlagRequired("source-server-id")
	})
	drsCmd.AddCommand(drs_updateReplicationConfigurationCmd)
}
