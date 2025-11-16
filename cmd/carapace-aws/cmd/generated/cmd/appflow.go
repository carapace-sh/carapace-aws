package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflowCmd = &cobra.Command{
	Use:   "appflow",
	Short: "Welcome to the Amazon AppFlow API reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflowCmd).Standalone()

	rootCmd.AddCommand(appflowCmd)
}
