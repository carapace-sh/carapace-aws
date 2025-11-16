package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(s3Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(s3Cmd).Standalone()

	})
	rootCmd.AddCommand(s3Cmd)
}
