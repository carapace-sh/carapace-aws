package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discoveryCmd = &cobra.Command{
	Use:   "discovery",
	Short: "Amazon Web Services Application Discovery Service\n\nAmazon Web Services Application Discovery Service (Application Discovery Service) helps you plan application migration projects.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discoveryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discoveryCmd).Standalone()

	})
	rootCmd.AddCommand(discoveryCmd)
}
