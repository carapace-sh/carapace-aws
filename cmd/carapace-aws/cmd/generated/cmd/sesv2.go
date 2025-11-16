package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2Cmd = &cobra.Command{
	Use:   "sesv2",
	Short: "Amazon SES API v2",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sesv2Cmd).Standalone()

	})
	rootCmd.AddCommand(sesv2Cmd)
}
