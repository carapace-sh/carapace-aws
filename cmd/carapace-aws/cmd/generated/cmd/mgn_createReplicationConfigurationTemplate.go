package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_createReplicationConfigurationTemplateCmd = &cobra.Command{
	Use:   "create-replication-configuration-template",
	Short: "Creates a new ReplicationConfigurationTemplate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_createReplicationConfigurationTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_createReplicationConfigurationTemplateCmd).Standalone()

		mgn_createReplicationConfigurationTemplateCmd.Flags().Bool("associate-default-security-group", false, "Request to associate the default Application Migration Service Security group with the Replication Settings template.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().String("bandwidth-throttling", "", "Request to configure bandwidth throttling during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().Bool("create-public-ip", false, "Request to create Public IP during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().String("data-plane-routing", "", "Request to configure data plane routing during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().String("default-large-staging-disk-type", "", "Request to configure the default large staging disk EBS volume type during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().String("ebs-encryption", "", "Request to configure EBS encryption during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().String("ebs-encryption-key-arn", "", "Request to configure an EBS encryption key during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().Bool("no-associate-default-security-group", false, "Request to associate the default Application Migration Service Security group with the Replication Settings template.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().Bool("no-create-public-ip", false, "Request to create Public IP during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().Bool("no-use-dedicated-replication-server", false, "Request to use Dedicated Replication Servers during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().Bool("no-use-fips-endpoint", false, "Request to use Fips Endpoint during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().String("replication-server-instance-type", "", "Request to configure the Replication Server instance type during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().String("replication-servers-security-groups-ids", "", "Request to configure the Replication Server Security group ID during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().String("staging-area-subnet-id", "", "Request to configure the Staging Area subnet ID during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().String("staging-area-tags", "", "Request to configure Staging Area tags during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().String("tags", "", "Request to configure tags during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().Bool("use-dedicated-replication-server", false, "Request to use Dedicated Replication Servers during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.Flags().Bool("use-fips-endpoint", false, "Request to use Fips Endpoint during Replication Settings template creation.")
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("associate-default-security-group")
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("bandwidth-throttling")
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("create-public-ip")
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("data-plane-routing")
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("default-large-staging-disk-type")
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("ebs-encryption")
		mgn_createReplicationConfigurationTemplateCmd.Flag("no-associate-default-security-group").Hidden = true
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("no-associate-default-security-group")
		mgn_createReplicationConfigurationTemplateCmd.Flag("no-create-public-ip").Hidden = true
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("no-create-public-ip")
		mgn_createReplicationConfigurationTemplateCmd.Flag("no-use-dedicated-replication-server").Hidden = true
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("no-use-dedicated-replication-server")
		mgn_createReplicationConfigurationTemplateCmd.Flag("no-use-fips-endpoint").Hidden = true
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("replication-server-instance-type")
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("replication-servers-security-groups-ids")
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("staging-area-subnet-id")
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("staging-area-tags")
		mgn_createReplicationConfigurationTemplateCmd.MarkFlagRequired("use-dedicated-replication-server")
	})
	mgnCmd.AddCommand(mgn_createReplicationConfigurationTemplateCmd)
}
