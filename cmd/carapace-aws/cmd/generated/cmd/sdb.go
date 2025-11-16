package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sdbCmd = &cobra.Command{
	Use:   "sdb",
	Short: "Amazon SimpleDB is a web service providing the core database functions of data indexing and querying in the cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sdbCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sdbCmd).Standalone()

	})
	rootCmd.AddCommand(sdbCmd)
}
