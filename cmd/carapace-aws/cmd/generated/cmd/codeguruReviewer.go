package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeguruReviewerCmd = &cobra.Command{
	Use:   "codeguru-reviewer",
	Short: "This section provides documentation for the Amazon CodeGuru Reviewer API operations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeguruReviewerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeguruReviewerCmd).Standalone()

	})
	rootCmd.AddCommand(codeguruReviewerCmd)
}
