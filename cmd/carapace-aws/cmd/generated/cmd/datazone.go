package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazoneCmd = &cobra.Command{
	Use:   "datazone",
	Short: "Amazon DataZone is a data management service that enables you to catalog, discover, govern, share, and analyze your data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazoneCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazoneCmd).Standalone()

	})
	rootCmd.AddCommand(datazoneCmd)
}
