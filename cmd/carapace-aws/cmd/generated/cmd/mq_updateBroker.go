package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mq_updateBrokerCmd = &cobra.Command{
	Use:   "update-broker",
	Short: "Adds a pending configuration change to a broker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mq_updateBrokerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mq_updateBrokerCmd).Standalone()

		mq_updateBrokerCmd.Flags().String("authentication-strategy", "", "Optional.")
		mq_updateBrokerCmd.Flags().String("auto-minor-version-upgrade", "", "Enables automatic upgrades to new patch versions for brokers as new versions are released and supported by Amazon MQ.")
		mq_updateBrokerCmd.Flags().String("broker-id", "", "The unique ID that Amazon MQ generates for the broker.")
		mq_updateBrokerCmd.Flags().String("configuration", "", "A list of information about the configuration.")
		mq_updateBrokerCmd.Flags().String("data-replication-mode", "", "Defines whether this broker is a part of a data replication pair.")
		mq_updateBrokerCmd.Flags().String("engine-version", "", "The broker engine version.")
		mq_updateBrokerCmd.Flags().String("host-instance-type", "", "The broker's host instance type to upgrade to.")
		mq_updateBrokerCmd.Flags().String("ldap-server-metadata", "", "Optional.")
		mq_updateBrokerCmd.Flags().String("logs", "", "Enables Amazon CloudWatch logging for brokers.")
		mq_updateBrokerCmd.Flags().String("maintenance-window-start-time", "", "The parameters that determine the WeeklyStartTime.")
		mq_updateBrokerCmd.Flags().String("security-groups", "", "The list of security groups (1 minimum, 5 maximum) that authorizes connections to brokers.")
		mq_updateBrokerCmd.MarkFlagRequired("broker-id")
	})
	mqCmd.AddCommand(mq_updateBrokerCmd)
}
