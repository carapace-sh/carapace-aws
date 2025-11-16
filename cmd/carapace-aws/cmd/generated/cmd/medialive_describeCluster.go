package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeClusterCmd = &cobra.Command{
	Use:   "describe-cluster",
	Short: "Get details about a Cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_describeClusterCmd).Standalone()

		medialive_describeClusterCmd.Flags().String("cluster-id", "", "The ID of the cluster.")
		medialive_describeClusterCmd.MarkFlagRequired("cluster-id")
	})
	medialiveCmd.AddCommand(medialive_describeClusterCmd)
}
