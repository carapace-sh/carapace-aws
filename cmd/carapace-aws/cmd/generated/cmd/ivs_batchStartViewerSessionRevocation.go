package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_batchStartViewerSessionRevocationCmd = &cobra.Command{
	Use:   "batch-start-viewer-session-revocation",
	Short: "Performs [StartViewerSessionRevocation]() on multiple channel ARN and viewer ID pairs simultaneously.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_batchStartViewerSessionRevocationCmd).Standalone()

	ivs_batchStartViewerSessionRevocationCmd.Flags().String("viewer-sessions", "", "Array of viewer sessions, one per channel-ARN and viewer-ID pair.")
	ivs_batchStartViewerSessionRevocationCmd.MarkFlagRequired("viewer-sessions")
	ivsCmd.AddCommand(ivs_batchStartViewerSessionRevocationCmd)
}
