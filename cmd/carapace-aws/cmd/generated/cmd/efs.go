package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var efsCmd = &cobra.Command{
	Use:   "efs",
	Short: "Amazon Elastic File System",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(efsCmd).Standalone()

	})
	rootCmd.AddCommand(efsCmd)
}
