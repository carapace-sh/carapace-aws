package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloud9Cmd = &cobra.Command{
	Use:   "cloud9",
	Short: "Cloud9\n\nCloud9 is a collection of tools that you can use to code, build, run, test, debug, and release software in the cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloud9Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloud9Cmd).Standalone()

	})
	rootCmd.AddCommand(cloud9Cmd)
}
