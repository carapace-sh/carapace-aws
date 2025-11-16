package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_putCaseEventConfigurationCmd = &cobra.Command{
	Use:   "put-case-event-configuration",
	Short: "Adds case event publishing configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_putCaseEventConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_putCaseEventConfigurationCmd).Standalone()

		connectcases_putCaseEventConfigurationCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_putCaseEventConfigurationCmd.Flags().String("event-bridge", "", "Configuration to enable EventBridge case event delivery and determine what data is delivered.")
		connectcases_putCaseEventConfigurationCmd.MarkFlagRequired("domain-id")
		connectcases_putCaseEventConfigurationCmd.MarkFlagRequired("event-bridge")
	})
	connectcasesCmd.AddCommand(connectcases_putCaseEventConfigurationCmd)
}
