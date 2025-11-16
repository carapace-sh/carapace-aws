package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appstream_deleteApplicationCmd).Standalone()

		appstream_deleteApplicationCmd.Flags().String("name", "", "The name of the application.")
		appstream_deleteApplicationCmd.MarkFlagRequired("name")
	})
	appstreamCmd.AddCommand(appstream_deleteApplicationCmd)
}
