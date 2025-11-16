package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ses_listConfigurationSetsCmd = &cobra.Command{
	Use:   "list-configuration-sets",
	Short: "Provides a list of the configuration sets associated with your Amazon SES account in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ses_listConfigurationSetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ses_listConfigurationSetsCmd).Standalone()

		ses_listConfigurationSetsCmd.Flags().String("max-items", "", "The number of configuration sets to return.")
		ses_listConfigurationSetsCmd.Flags().String("next-token", "", "A token returned from a previous call to `ListConfigurationSets` to indicate the position of the configuration set in the configuration set list.")
	})
	sesCmd.AddCommand(ses_listConfigurationSetsCmd)
}
