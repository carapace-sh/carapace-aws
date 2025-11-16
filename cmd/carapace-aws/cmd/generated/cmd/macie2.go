package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2Cmd = &cobra.Command{
	Use:   "macie2",
	Short: "Amazon Macie",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2Cmd).Standalone()

	rootCmd.AddCommand(macie2Cmd)
}
