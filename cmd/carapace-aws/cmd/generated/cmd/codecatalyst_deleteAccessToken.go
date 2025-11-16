package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_deleteAccessTokenCmd = &cobra.Command{
	Use:   "delete-access-token",
	Short: "Deletes a specified personal access token (PAT).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_deleteAccessTokenCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_deleteAccessTokenCmd).Standalone()

		codecatalyst_deleteAccessTokenCmd.Flags().String("id", "", "The ID of the personal access token to delete.")
		codecatalyst_deleteAccessTokenCmd.MarkFlagRequired("id")
	})
	codecatalystCmd.AddCommand(codecatalyst_deleteAccessTokenCmd)
}
