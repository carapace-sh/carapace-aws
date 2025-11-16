package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_createBrokerCmd = &cobra.Command{
	Use:   "create-broker",
	Short: "Creates a broker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_createBrokerCmd).Standalone()

	mq_createBrokerCmd.Flags().String("authentication-strategy", "", "Optional.")
	mq_createBrokerCmd.Flags().String("auto-minor-version-upgrade", "", "Enables automatic upgrades to new patch versions for brokers as new versions are released and supported by Amazon MQ.")
	mq_createBrokerCmd.Flags().String("broker-name", "", "Required.")
	mq_createBrokerCmd.Flags().String("configuration", "", "A list of information about the configuration.")
	mq_createBrokerCmd.Flags().String("creator-request-id", "", "The unique ID that the requester receives for the created broker.")
	mq_createBrokerCmd.Flags().String("data-replication-mode", "", "Defines whether this broker is a part of a data replication pair.")
	mq_createBrokerCmd.Flags().String("data-replication-primary-broker-arn", "", "The Amazon Resource Name (ARN) of the primary broker that is used to replicate data from in a data replication pair, and is applied to the replica broker.")
	mq_createBrokerCmd.Flags().String("deployment-mode", "", "Required.")
	mq_createBrokerCmd.Flags().String("encryption-options", "", "Encryption options for the broker.")
	mq_createBrokerCmd.Flags().String("engine-type", "", "Required.")
	mq_createBrokerCmd.Flags().String("engine-version", "", "The broker engine version.")
	mq_createBrokerCmd.Flags().String("host-instance-type", "", "Required.")
	mq_createBrokerCmd.Flags().String("ldap-server-metadata", "", "Optional.")
	mq_createBrokerCmd.Flags().String("logs", "", "Enables Amazon CloudWatch logging for brokers.")
	mq_createBrokerCmd.Flags().String("maintenance-window-start-time", "", "The parameters that determine the WeeklyStartTime.")
	mq_createBrokerCmd.Flags().String("publicly-accessible", "", "Enables connections from applications outside of the VPC that hosts the broker's subnets.")
	mq_createBrokerCmd.Flags().String("security-groups", "", "The list of rules (1 minimum, 125 maximum) that authorize connections to brokers.")
	mq_createBrokerCmd.Flags().String("storage-type", "", "The broker's storage type.")
	mq_createBrokerCmd.Flags().String("subnet-ids", "", "The list of groups that define which subnets and IP ranges the broker can use from different Availability Zones.")
	mq_createBrokerCmd.Flags().String("tags", "", "Create tags when creating the broker.")
	mq_createBrokerCmd.Flags().String("users", "", "The list of broker users (persons or applications) who can access queues and topics.")
	mq_createBrokerCmd.MarkFlagRequired("broker-name")
	mq_createBrokerCmd.MarkFlagRequired("deployment-mode")
	mq_createBrokerCmd.MarkFlagRequired("engine-type")
	mq_createBrokerCmd.MarkFlagRequired("host-instance-type")
	mq_createBrokerCmd.MarkFlagRequired("publicly-accessible")
	mqCmd.AddCommand(mq_createBrokerCmd)
}
