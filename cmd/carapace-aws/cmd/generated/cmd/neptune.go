package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneCmd = &cobra.Command{
	Use:   "neptune",
	Short: "Amazon Neptune",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneCmd).Standalone()

	})
	rootCmd.AddCommand(neptuneCmd)
}
