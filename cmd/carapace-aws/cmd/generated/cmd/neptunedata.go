package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedataCmd = &cobra.Command{
	Use:   "neptunedata",
	Short: "Neptune Data API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptunedataCmd).Standalone()

	})
	rootCmd.AddCommand(neptunedataCmd)
}
