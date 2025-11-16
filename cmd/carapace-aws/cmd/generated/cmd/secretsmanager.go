package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secretsmanagerCmd = &cobra.Command{
	Use:   "secretsmanager",
	Short: "Amazon Web Services Secrets Manager\n\nAmazon Web Services Secrets Manager provides a service to enable you to store, manage, and retrieve, secrets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secretsmanagerCmd).Standalone()

	rootCmd.AddCommand(secretsmanagerCmd)
}
