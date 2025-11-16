package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workmailCmd = &cobra.Command{
	Use:   "workmail",
	Short: "WorkMail is a secure, managed business email and calendaring service with support for existing desktop and mobile email clients.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workmailCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workmailCmd).Standalone()

	})
	rootCmd.AddCommand(workmailCmd)
}
