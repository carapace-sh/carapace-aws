package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connectcases_getCaseEventConfigurationCmd = &cobra.Command{
	Use:   "get-case-event-configuration",
	Short: "Returns the case event publishing configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connectcases_getCaseEventConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connectcases_getCaseEventConfigurationCmd).Standalone()

		connectcases_getCaseEventConfigurationCmd.Flags().String("domain-id", "", "The unique identifier of the Cases domain.")
		connectcases_getCaseEventConfigurationCmd.MarkFlagRequired("domain-id")
	})
	connectcasesCmd.AddCommand(connectcases_getCaseEventConfigurationCmd)
}
