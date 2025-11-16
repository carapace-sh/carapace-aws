package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasyncCmd = &cobra.Command{
	Use:   "datasync",
	Short: "DataSync",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasyncCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datasyncCmd).Standalone()

	})
	rootCmd.AddCommand(datasyncCmd)
}
