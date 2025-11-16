package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elbCmd = &cobra.Command{
	Use:   "elb",
	Short: "Elastic Load Balancing",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elbCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elbCmd).Standalone()

	})
	rootCmd.AddCommand(elbCmd)
}
