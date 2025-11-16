package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_describeClustersCmd = &cobra.Command{
	Use:   "describe-clusters",
	Short: "Describes one or more of your clusters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_describeClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_describeClustersCmd).Standalone()

		ecs_describeClustersCmd.Flags().String("clusters", "", "A list of up to 100 cluster names or full cluster Amazon Resource Name (ARN) entries.")
		ecs_describeClustersCmd.Flags().String("include", "", "Determines whether to include additional information about the clusters in the response.")
	})
	ecsCmd.AddCommand(ecs_describeClustersCmd)
}
