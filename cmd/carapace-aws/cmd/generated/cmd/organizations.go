package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var organizationsCmd = &cobra.Command{
	Use:   "organizations",
	Short: "Organizations is a web service that enables you to consolidate your multiple Amazon Web Services accounts into an *organization* and centrally manage your accounts and their resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(organizationsCmd).Standalone()

	rootCmd.AddCommand(organizationsCmd)
}
