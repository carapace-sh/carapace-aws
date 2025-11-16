package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_updateReplicationConfigurationCmd = &cobra.Command{
	Use:   "update-replication-configuration",
	Short: "Allows you to update multiple ReplicationConfigurations by Source Server ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_updateReplicationConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_updateReplicationConfigurationCmd).Standalone()

		mgn_updateReplicationConfigurationCmd.Flags().String("account-id", "", "Update replication configuration Account ID request.")
		mgn_updateReplicationConfigurationCmd.Flags().Bool("associate-default-security-group", false, "Update replication configuration associate default Application Migration Service Security group request.")
		mgn_updateReplicationConfigurationCmd.Flags().String("bandwidth-throttling", "", "Update replication configuration bandwidth throttling request.")
		mgn_updateReplicationConfigurationCmd.Flags().Bool("create-public-ip", false, "Update replication configuration create Public IP request.")
		mgn_updateReplicationConfigurationCmd.Flags().String("data-plane-routing", "", "Update replication configuration data plane routing request.")
		mgn_updateReplicationConfigurationCmd.Flags().String("default-large-staging-disk-type", "", "Update replication configuration use default large Staging Disk type request.")
		mgn_updateReplicationConfigurationCmd.Flags().String("ebs-encryption", "", "Update replication configuration EBS encryption request.")
		mgn_updateReplicationConfigurationCmd.Flags().String("ebs-encryption-key-arn", "", "Update replication configuration EBS encryption key ARN request.")
		mgn_updateReplicationConfigurationCmd.Flags().String("name", "", "Update replication configuration name request.")
		mgn_updateReplicationConfigurationCmd.Flags().Bool("no-associate-default-security-group", false, "Update replication configuration associate default Application Migration Service Security group request.")
		mgn_updateReplicationConfigurationCmd.Flags().Bool("no-create-public-ip", false, "Update replication configuration create Public IP request.")
		mgn_updateReplicationConfigurationCmd.Flags().Bool("no-use-dedicated-replication-server", false, "Update replication configuration use dedicated Replication Server request.")
		mgn_updateReplicationConfigurationCmd.Flags().Bool("no-use-fips-endpoint", false, "Update replication configuration use Fips Endpoint.")
		mgn_updateReplicationConfigurationCmd.Flags().String("replicated-disks", "", "Update replication configuration replicated disks request.")
		mgn_updateReplicationConfigurationCmd.Flags().String("replication-server-instance-type", "", "Update replication configuration Replication Server instance type request.")
		mgn_updateReplicationConfigurationCmd.Flags().String("replication-servers-security-groups-ids", "", "Update replication configuration Replication Server Security Groups IDs request.")
		mgn_updateReplicationConfigurationCmd.Flags().String("source-server-id", "", "Update replication configuration Source Server ID request.")
		mgn_updateReplicationConfigurationCmd.Flags().String("staging-area-subnet-id", "", "Update replication configuration Staging Area subnet request.")
		mgn_updateReplicationConfigurationCmd.Flags().String("staging-area-tags", "", "Update replication configuration Staging Area Tags request.")
		mgn_updateReplicationConfigurationCmd.Flags().Bool("use-dedicated-replication-server", false, "Update replication configuration use dedicated Replication Server request.")
		mgn_updateReplicationConfigurationCmd.Flags().Bool("use-fips-endpoint", false, "Update replication configuration use Fips Endpoint.")
		mgn_updateReplicationConfigurationCmd.Flag("no-associate-default-security-group").Hidden = true
		mgn_updateReplicationConfigurationCmd.Flag("no-create-public-ip").Hidden = true
		mgn_updateReplicationConfigurationCmd.Flag("no-use-dedicated-replication-server").Hidden = true
		mgn_updateReplicationConfigurationCmd.Flag("no-use-fips-endpoint").Hidden = true
		mgn_updateReplicationConfigurationCmd.MarkFlagRequired("source-server-id")
	})
	mgnCmd.AddCommand(mgn_updateReplicationConfigurationCmd)
}
