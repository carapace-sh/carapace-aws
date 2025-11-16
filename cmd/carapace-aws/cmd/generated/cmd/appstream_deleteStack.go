package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_deleteStackCmd = &cobra.Command{
	Use:   "delete-stack",
	Short: "Deletes the specified stack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_deleteStackCmd).Standalone()

	appstream_deleteStackCmd.Flags().String("name", "", "The name of the stack.")
	appstream_deleteStackCmd.MarkFlagRequired("name")
	appstreamCmd.AddCommand(appstream_deleteStackCmd)
}
