package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appstream_expireSessionCmd = &cobra.Command{
	Use:   "expire-session",
	Short: "Immediately stops the specified streaming session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appstream_expireSessionCmd).Standalone()

	appstream_expireSessionCmd.Flags().String("session-id", "", "The identifier of the streaming session.")
	appstream_expireSessionCmd.MarkFlagRequired("session-id")
	appstreamCmd.AddCommand(appstream_expireSessionCmd)
}
