package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_listSecurityConfigurationsCmd = &cobra.Command{
	Use:   "list-security-configurations",
	Short: "Lists security configurations based on a set of parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_listSecurityConfigurationsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emrContainers_listSecurityConfigurationsCmd).Standalone()

		emrContainers_listSecurityConfigurationsCmd.Flags().String("created-after", "", "The date and time after which the security configuration was created.")
		emrContainers_listSecurityConfigurationsCmd.Flags().String("created-before", "", "The date and time before which the security configuration was created.")
		emrContainers_listSecurityConfigurationsCmd.Flags().String("max-results", "", "The maximum number of security configurations the operation can list.")
		emrContainers_listSecurityConfigurationsCmd.Flags().String("next-token", "", "The token for the next set of security configurations to return.")
	})
	emrContainersCmd.AddCommand(emrContainers_listSecurityConfigurationsCmd)
}
