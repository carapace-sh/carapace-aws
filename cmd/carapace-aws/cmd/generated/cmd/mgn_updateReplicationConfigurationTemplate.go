package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_updateReplicationConfigurationTemplateCmd = &cobra.Command{
	Use:   "update-replication-configuration-template",
	Short: "Updates multiple ReplicationConfigurationTemplates by ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_updateReplicationConfigurationTemplateCmd).Standalone()

	mgn_updateReplicationConfigurationTemplateCmd.Flags().String("arn", "", "Update replication configuration template ARN request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().Bool("associate-default-security-group", false, "Update replication configuration template associate default Application Migration Service Security group request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().String("bandwidth-throttling", "", "Update replication configuration template bandwidth throttling request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().Bool("create-public-ip", false, "Update replication configuration template create Public IP request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().String("data-plane-routing", "", "Update replication configuration template data plane routing request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().String("default-large-staging-disk-type", "", "Update replication configuration template use default large Staging Disk type request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().String("ebs-encryption", "", "Update replication configuration template EBS encryption request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().String("ebs-encryption-key-arn", "", "Update replication configuration template EBS encryption key ARN request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().Bool("no-associate-default-security-group", false, "Update replication configuration template associate default Application Migration Service Security group request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().Bool("no-create-public-ip", false, "Update replication configuration template create Public IP request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().Bool("no-use-dedicated-replication-server", false, "Update replication configuration template use dedicated Replication Server request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().Bool("no-use-fips-endpoint", false, "Update replication configuration template use Fips Endpoint request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().String("replication-configuration-template-id", "", "Update replication configuration template template ID request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().String("replication-server-instance-type", "", "Update replication configuration template Replication Server instance type request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().String("replication-servers-security-groups-ids", "", "Update replication configuration template Replication Server Security groups IDs request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().String("staging-area-subnet-id", "", "Update replication configuration template Staging Area subnet ID request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().String("staging-area-tags", "", "Update replication configuration template Staging Area Tags request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().Bool("use-dedicated-replication-server", false, "Update replication configuration template use dedicated Replication Server request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flags().Bool("use-fips-endpoint", false, "Update replication configuration template use Fips Endpoint request.")
	mgn_updateReplicationConfigurationTemplateCmd.Flag("no-associate-default-security-group").Hidden = true
	mgn_updateReplicationConfigurationTemplateCmd.Flag("no-create-public-ip").Hidden = true
	mgn_updateReplicationConfigurationTemplateCmd.Flag("no-use-dedicated-replication-server").Hidden = true
	mgn_updateReplicationConfigurationTemplateCmd.Flag("no-use-fips-endpoint").Hidden = true
	mgn_updateReplicationConfigurationTemplateCmd.MarkFlagRequired("replication-configuration-template-id")
	mgnCmd.AddCommand(mgn_updateReplicationConfigurationTemplateCmd)
}
