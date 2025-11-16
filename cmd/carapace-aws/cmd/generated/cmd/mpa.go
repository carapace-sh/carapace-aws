package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mpaCmd = &cobra.Command{
	Use:   "mpa",
	Short: "Multi-party approval is a capability of [Organizations](http://aws.amazon.com/organizations) that allows you to protect a predefined list of operations through a distributed approval process.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mpaCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mpaCmd).Standalone()

	})
	rootCmd.AddCommand(mpaCmd)
}
