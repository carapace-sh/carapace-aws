package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_startViewerSessionRevocationCmd = &cobra.Command{
	Use:   "start-viewer-session-revocation",
	Short: "Starts the process of revoking the viewer session associated with a specified channel ARN and viewer ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_startViewerSessionRevocationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ivs_startViewerSessionRevocationCmd).Standalone()

		ivs_startViewerSessionRevocationCmd.Flags().String("channel-arn", "", "The ARN of the channel associated with the viewer session to revoke.")
		ivs_startViewerSessionRevocationCmd.Flags().String("viewer-id", "", "The ID of the viewer associated with the viewer session to revoke.")
		ivs_startViewerSessionRevocationCmd.Flags().String("viewer-session-versions-less-than-or-equal-to", "", "An optional filter on which versions of the viewer session to revoke.")
		ivs_startViewerSessionRevocationCmd.MarkFlagRequired("channel-arn")
		ivs_startViewerSessionRevocationCmd.MarkFlagRequired("viewer-id")
	})
	ivsCmd.AddCommand(ivs_startViewerSessionRevocationCmd)
}
