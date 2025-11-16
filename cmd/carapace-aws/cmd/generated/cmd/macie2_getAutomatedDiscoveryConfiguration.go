package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getAutomatedDiscoveryConfigurationCmd = &cobra.Command{
	Use:   "get-automated-discovery-configuration",
	Short: "Retrieves the configuration settings and status of automated sensitive data discovery for an organization or standalone account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getAutomatedDiscoveryConfigurationCmd).Standalone()

	macie2Cmd.AddCommand(macie2_getAutomatedDiscoveryConfigurationCmd)
}
