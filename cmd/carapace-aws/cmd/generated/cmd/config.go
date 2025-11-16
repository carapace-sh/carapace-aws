package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Config\n\nConfig provides a way to keep track of the configurations of all the Amazon Web Services resources associated with your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(configCmd).Standalone()

	})
	rootCmd.AddCommand(configCmd)
}
