package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanagerCmd = &cobra.Command{
	Use:   "secretsmanager",
	Short: "Amazon Web Services Secrets Manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanagerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(secretsmanagerCmd).Standalone()

	})
	rootCmd.AddCommand(secretsmanagerCmd)
}
