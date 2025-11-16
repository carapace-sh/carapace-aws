package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarmCmd = &cobra.Command{
	Use:   "devicefarm",
	Short: "Welcome to the AWS Device Farm API documentation, which contains APIs for:",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarmCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarmCmd).Standalone()

	})
	rootCmd.AddCommand(devicefarmCmd)
}
