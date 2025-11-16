package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplifyuibuilder_refreshTokenCmd = &cobra.Command{
	Use:   "refresh-token",
	Short: "This is for internal use.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplifyuibuilder_refreshTokenCmd).Standalone()

	amplifyuibuilder_refreshTokenCmd.Flags().String("provider", "", "The third-party provider for the token.")
	amplifyuibuilder_refreshTokenCmd.Flags().String("refresh-token-body", "", "Information about the refresh token request.")
	amplifyuibuilder_refreshTokenCmd.MarkFlagRequired("provider")
	amplifyuibuilder_refreshTokenCmd.MarkFlagRequired("refresh-token-body")
	amplifyuibuilderCmd.AddCommand(amplifyuibuilder_refreshTokenCmd)
}
