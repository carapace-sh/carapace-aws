package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putConfigurationSetReputationOptionsCmd = &cobra.Command{
	Use:   "put-configuration-set-reputation-options",
	Short: "Enable or disable collection of reputation metrics for emails that you send using a particular configuration set in a specific Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putConfigurationSetReputationOptionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_putConfigurationSetReputationOptionsCmd).Standalone()

		sesv2_putConfigurationSetReputationOptionsCmd.Flags().String("configuration-set-name", "", "The name of the configuration set.")
		sesv2_putConfigurationSetReputationOptionsCmd.Flags().String("reputation-metrics-enabled", "", "If `true`, tracking of reputation metrics is enabled for the configuration set.")
		sesv2_putConfigurationSetReputationOptionsCmd.MarkFlagRequired("configuration-set-name")
	})
	sesv2Cmd.AddCommand(sesv2_putConfigurationSetReputationOptionsCmd)
}
