package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscoveryCmd = &cobra.Command{
	Use:   "servicediscovery",
	Short: "Cloud Map",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscoveryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscoveryCmd).Standalone()

	})
	rootCmd.AddCommand(servicediscoveryCmd)
}
