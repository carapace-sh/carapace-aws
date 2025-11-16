package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var daxCmd = &cobra.Command{
	Use:   "dax",
	Short: "DAX is a managed caching service engineered for Amazon DynamoDB.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(daxCmd).Standalone()

	rootCmd.AddCommand(daxCmd)
}
