package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_updateAutomatedDiscoveryConfigurationCmd = &cobra.Command{
	Use:   "update-automated-discovery-configuration",
	Short: "Changes the configuration settings and status of automated sensitive data discovery for an organization or standalone account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_updateAutomatedDiscoveryConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(macie2_updateAutomatedDiscoveryConfigurationCmd).Standalone()

		macie2_updateAutomatedDiscoveryConfigurationCmd.Flags().String("auto-enable-organization-members", "", "Specifies whether to automatically enable automated sensitive data discovery for accounts in the organization.")
		macie2_updateAutomatedDiscoveryConfigurationCmd.Flags().String("status", "", "The new status of automated sensitive data discovery for the organization or account.")
		macie2_updateAutomatedDiscoveryConfigurationCmd.MarkFlagRequired("status")
	})
	macie2Cmd.AddCommand(macie2_updateAutomatedDiscoveryConfigurationCmd)
}
