package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassCmd = &cobra.Command{
	Use:   "greengrass",
	Short: "AWS IoT Greengrass seamlessly extends AWS onto physical devices so they can act locally on the data they generate, while still using the cloud for management, analytics, and durable storage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassCmd).Standalone()

	})
	rootCmd.AddCommand(greengrassCmd)
}
