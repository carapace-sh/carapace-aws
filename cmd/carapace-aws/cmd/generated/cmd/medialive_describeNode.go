package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeNodeCmd = &cobra.Command{
	Use:   "describe-node",
	Short: "Get details about a Node in the specified Cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeNodeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_describeNodeCmd).Standalone()

		medialive_describeNodeCmd.Flags().String("cluster-id", "", "The ID of the cluster")
		medialive_describeNodeCmd.Flags().String("node-id", "", "The ID of the node.")
		medialive_describeNodeCmd.MarkFlagRequired("cluster-id")
		medialive_describeNodeCmd.MarkFlagRequired("node-id")
	})
	medialiveCmd.AddCommand(medialive_describeNodeCmd)
}
