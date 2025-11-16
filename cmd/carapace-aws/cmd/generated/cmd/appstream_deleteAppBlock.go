package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteAppBlockCmd = &cobra.Command{
	Use:   "delete-app-block",
	Short: "Deletes an app block.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteAppBlockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_deleteAppBlockCmd).Standalone()

		appstream_deleteAppBlockCmd.Flags().String("name", "", "The name of the app block.")
		appstream_deleteAppBlockCmd.MarkFlagRequired("name")
	})
	appstreamCmd.AddCommand(appstream_deleteAppBlockCmd)
}
