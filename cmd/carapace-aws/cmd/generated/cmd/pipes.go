package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipesCmd = &cobra.Command{
	Use:   "pipes",
	Short: "Amazon EventBridge Pipes connects event sources to targets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipesCmd).Standalone()

	rootCmd.AddCommand(pipesCmd)
}
