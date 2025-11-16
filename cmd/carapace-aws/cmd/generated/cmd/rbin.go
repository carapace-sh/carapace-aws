package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbinCmd = &cobra.Command{
	Use:   "rbin",
	Short: "This is the *Recycle Bin API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbinCmd).Standalone()

	rootCmd.AddCommand(rbinCmd)
}
