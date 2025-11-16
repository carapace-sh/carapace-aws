package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfigdataCmd = &cobra.Command{
	Use:   "appconfigdata",
	Short: "AppConfig Data provides the data plane APIs your application uses to retrieve configuration data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfigdataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfigdataCmd).Standalone()

	})
	rootCmd.AddCommand(appconfigdataCmd)
}
