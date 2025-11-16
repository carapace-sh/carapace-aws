package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnectionsCmd = &cobra.Command{
	Use:   "codeconnections",
	Short: "AWS CodeConnections",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeconnectionsCmd).Standalone()

	})
	rootCmd.AddCommand(codeconnectionsCmd)
}
