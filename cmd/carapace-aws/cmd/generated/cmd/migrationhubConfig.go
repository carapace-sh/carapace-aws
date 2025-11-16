package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var migrationhubConfigCmd = &cobra.Command{
	Use:   "migrationhub-config",
	Short: "The AWS Migration Hub home region APIs are available specifically for working with your Migration Hub home region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(migrationhubConfigCmd).Standalone()

	rootCmd.AddCommand(migrationhubConfigCmd)
}
