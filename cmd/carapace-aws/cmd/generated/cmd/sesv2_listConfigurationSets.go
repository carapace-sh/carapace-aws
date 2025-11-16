package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_listConfigurationSetsCmd = &cobra.Command{
	Use:   "list-configuration-sets",
	Short: "List all of the configuration sets associated with your account in the current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_listConfigurationSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2_listConfigurationSetsCmd).Standalone()

		sesv2_listConfigurationSetsCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListConfigurationSets` to indicate the position in the list of configuration sets.")
		sesv2_listConfigurationSetsCmd.Flags().String("page-size", "", "The number of results to show in a single call to `ListConfigurationSets`.")
	})
	sesv2Cmd.AddCommand(sesv2_listConfigurationSetsCmd)
}
