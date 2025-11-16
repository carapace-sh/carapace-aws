package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_createAccessTokenCmd = &cobra.Command{
	Use:   "create-access-token",
	Short: "Creates a personal access token (PAT) for the current user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_createAccessTokenCmd).Standalone()

	codecatalyst_createAccessTokenCmd.Flags().String("expires-time", "", "The date and time the personal access token expires, in coordinated universal time (UTC) timestamp format as specified in [RFC 3339](https://www.rfc-editor.org/rfc/rfc3339#section-5.6).")
	codecatalyst_createAccessTokenCmd.Flags().String("name", "", "The friendly name of the personal access token.")
	codecatalyst_createAccessTokenCmd.MarkFlagRequired("name")
	codecatalystCmd.AddCommand(codecatalyst_createAccessTokenCmd)
}
