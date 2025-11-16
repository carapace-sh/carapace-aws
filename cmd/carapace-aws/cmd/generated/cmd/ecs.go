package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecsCmd = &cobra.Command{
	Use:   "ecs",
	Short: "Amazon Elastic Container Service",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecsCmd).Standalone()

	})
	rootCmd.AddCommand(ecsCmd)
}
