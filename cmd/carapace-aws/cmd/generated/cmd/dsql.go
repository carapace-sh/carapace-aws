package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dsqlCmd = &cobra.Command{
	Use:   "dsql",
	Short: "This is an interface reference for Amazon Aurora DSQL.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dsqlCmd).Standalone()

	rootCmd.AddCommand(dsqlCmd)
}
