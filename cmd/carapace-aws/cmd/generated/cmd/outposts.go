package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outpostsCmd = &cobra.Command{
	Use:   "outposts",
	Short: "Amazon Web Services Outposts is a fully managed service that extends Amazon Web Services infrastructure, APIs, and tools to customer premises.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outpostsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outpostsCmd).Standalone()

	})
	rootCmd.AddCommand(outpostsCmd)
}
