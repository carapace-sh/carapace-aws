package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarmCmd = &cobra.Command{
	Use:   "devicefarm",
	Short: "Welcome to the AWS Device Farm API documentation, which contains APIs for:\n\n- Testing on desktop browsers\n  \n  Device Farm makes it possible for you to test your web applications on desktop browsers using Selenium.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarmCmd).Standalone()

	})
	rootCmd.AddCommand(devicefarmCmd)
}
