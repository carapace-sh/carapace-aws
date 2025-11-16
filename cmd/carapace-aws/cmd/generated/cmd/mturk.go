package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturkCmd = &cobra.Command{
	Use:   "mturk",
	Short: "Amazon Mechanical Turk API Reference",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturkCmd).Standalone()

	})
	rootCmd.AddCommand(mturkCmd)
}
