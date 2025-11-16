package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_exchangeCodeForTokenCmd = &cobra.Command{
	Use:   "exchange-code-for-token",
	Short: "This is for internal use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_exchangeCodeForTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amplifyuibuilder_exchangeCodeForTokenCmd).Standalone()

		amplifyuibuilder_exchangeCodeForTokenCmd.Flags().String("provider", "", "The third-party provider for the token.")
		amplifyuibuilder_exchangeCodeForTokenCmd.Flags().String("request", "", "Describes the configuration of the request.")
		amplifyuibuilder_exchangeCodeForTokenCmd.MarkFlagRequired("provider")
		amplifyuibuilder_exchangeCodeForTokenCmd.MarkFlagRequired("request")
	})
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_exchangeCodeForTokenCmd)
}
