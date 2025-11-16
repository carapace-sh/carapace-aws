package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dmsCmd = &cobra.Command{
	Use:   "dms",
	Short: "Database Migration Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dmsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dmsCmd).Standalone()

	})
	rootCmd.AddCommand(dmsCmd)
}
