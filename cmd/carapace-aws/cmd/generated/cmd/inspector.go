package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspectorCmd = &cobra.Command{
	Use:   "inspector",
	Short: "Amazon Inspector",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspectorCmd).Standalone()

	})
	rootCmd.AddCommand(inspectorCmd)
}
