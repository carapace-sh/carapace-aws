package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhubCmd = &cobra.Command{
	Use:   "securityhub",
	Short: "Security Hub provides you with a comprehensive view of your security state in Amazon Web Services and helps you assess your Amazon Web Services environment against security industry standards and best practices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhubCmd).Standalone()

	rootCmd.AddCommand(securityhubCmd)
}
