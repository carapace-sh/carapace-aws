package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_updateConfigurationSetReputationMetricsEnabledCmd = &cobra.Command{
	Use:   "update-configuration-set-reputation-metrics-enabled",
	Short: "Enables or disables the publishing of reputation metrics for emails sent using a specific configuration set in a given Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_updateConfigurationSetReputationMetricsEnabledCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_updateConfigurationSetReputationMetricsEnabledCmd).Standalone()

		ses_updateConfigurationSetReputationMetricsEnabledCmd.Flags().String("configuration-set-name", "", "The name of the configuration set to update.")
		ses_updateConfigurationSetReputationMetricsEnabledCmd.Flags().String("enabled", "", "Describes whether or not Amazon SES publishes reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch.")
		ses_updateConfigurationSetReputationMetricsEnabledCmd.MarkFlagRequired("configuration-set-name")
		ses_updateConfigurationSetReputationMetricsEnabledCmd.MarkFlagRequired("enabled")
	})
	sesCmd.AddCommand(ses_updateConfigurationSetReputationMetricsEnabledCmd)
}
