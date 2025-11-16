package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsailCmd = &cobra.Command{
	Use:   "lightsail",
	Short: "Amazon Lightsail is the easiest way to get started with Amazon Web Services (Amazon Web Services) for developers who need to build websites or web applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsailCmd).Standalone()

	})
	rootCmd.AddCommand(lightsailCmd)
}
