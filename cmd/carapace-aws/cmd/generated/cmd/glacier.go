package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glacierCmd = &cobra.Command{
	Use:   "glacier",
	Short: "Amazon S3 Glacier (Glacier) is a storage solution for \"cold data.\"",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glacierCmd).Standalone()

	rootCmd.AddCommand(glacierCmd)
}
