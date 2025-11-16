package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mailmanagerCmd = &cobra.Command{
	Use:   "mailmanager",
	Short: "Amazon SES Mail Manager API",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mailmanagerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mailmanagerCmd).Standalone()

	})
	rootCmd.AddCommand(mailmanagerCmd)
}
