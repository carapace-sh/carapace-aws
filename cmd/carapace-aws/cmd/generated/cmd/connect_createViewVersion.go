package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_createViewVersionCmd = &cobra.Command{
	Use:   "create-view-version",
	Short: "Publishes a new version of the view identifier.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_createViewVersionCmd).Standalone()

	connect_createViewVersionCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_createViewVersionCmd.Flags().String("version-description", "", "The description for the version being published.")
	connect_createViewVersionCmd.Flags().String("view-content-sha256", "", "Indicates the checksum value of the latest published view content.")
	connect_createViewVersionCmd.Flags().String("view-id", "", "The identifier of the view.")
	connect_createViewVersionCmd.MarkFlagRequired("instance-id")
	connect_createViewVersionCmd.MarkFlagRequired("view-id")
	connectCmd.AddCommand(connect_createViewVersionCmd)
}
