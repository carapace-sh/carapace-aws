package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfigdata_getLatestConfigurationCmd = &cobra.Command{
	Use:   "get-latest-configuration",
	Short: "Retrieves the latest deployed configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfigdata_getLatestConfigurationCmd).Standalone()

	appconfigdata_getLatestConfigurationCmd.Flags().String("configuration-token", "", "Token describing the current state of the configuration session.")
	appconfigdata_getLatestConfigurationCmd.MarkFlagRequired("configuration-token")
	appconfigdataCmd.AddCommand(appconfigdata_getLatestConfigurationCmd)
}
