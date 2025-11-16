package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsCmd = &cobra.Command{
	Use:   "ds",
	Short: "Directory Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dsCmd).Standalone()

	})
	rootCmd.AddCommand(dsCmd)
}
