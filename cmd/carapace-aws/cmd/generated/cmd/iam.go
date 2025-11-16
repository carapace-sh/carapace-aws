package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iamCmd = &cobra.Command{
	Use:   "iam",
	Short: "Identity and Access Management\n\nIdentity and Access Management (IAM) is a web service for securely controlling access to Amazon Web Services services.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iamCmd).Standalone()

	rootCmd.AddCommand(iamCmd)
}
