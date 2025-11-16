package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pinpointEmail_putConfigurationSetReputationOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-reputation-options",
	Short: "Enable or disable collection of reputation metrics for emails that you send using a particular configuration set in a specific AWS Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pinpointEmail_putConfigurationSetReputationOptionsCmd).Standalone()

	pinpointEmail_putConfigurationSetReputationOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set that you want to enable or disable reputation metric tracking for.")
	pinpointEmail_putConfigurationSetReputationOptionsCmd.Flags().String("reputation-metrics-enabled", "", "If `true`, tracking of reputation metrics is enabled for the configuration set.")
	pinpointEmail_putConfigurationSetReputationOptionsCmd.MarkFlagRequired("configuration-set-name")
	pinpointEmailCmd.AddCommand(pinpointEmail_putConfigurationSetReputationOptionsCmd)
}
