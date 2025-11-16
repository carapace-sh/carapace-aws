package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesCmd = &cobra.Command{
	Use:   "ses",
	Short: "Amazon Simple Email Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesCmd).Standalone()

	})
	rootCmd.AddCommand(sesCmd)
}
