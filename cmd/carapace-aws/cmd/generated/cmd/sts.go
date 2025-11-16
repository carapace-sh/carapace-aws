package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stsCmd = &cobra.Command{
	Use:   "sts",
	Short: "Security Token Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(stsCmd).Standalone()

	})
	rootCmd.AddCommand(stsCmd)
}
