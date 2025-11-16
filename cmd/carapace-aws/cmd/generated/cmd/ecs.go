package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecsCmd = &cobra.Command{
	Use:   "ecs",
	Short: "Amazon Elastic Container Service\n\nAmazon Elastic Container Service (Amazon ECS) is a highly scalable, fast, container management service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecsCmd).Standalone()

	rootCmd.AddCommand(ecsCmd)
}
