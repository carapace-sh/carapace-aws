package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsxCmd = &cobra.Command{
	Use:   "fsx",
	Short: "Amazon FSx is a fully managed service that makes it easy for storage and application administrators to launch and use shared file storage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsxCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fsxCmd).Standalone()

	})
	rootCmd.AddCommand(fsxCmd)
}
