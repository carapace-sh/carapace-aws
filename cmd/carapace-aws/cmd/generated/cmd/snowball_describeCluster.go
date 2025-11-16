package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_describeClusterCmd = &cobra.Command{
	Use:   "describe-cluster",
	Short: "Returns information about a specific cluster including shipping information, cluster status, and other important metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_describeClusterCmd).Standalone()

	snowball_describeClusterCmd.Flags().String("cluster-id", "", "The automatically generated ID for a cluster.")
	snowball_describeClusterCmd.MarkFlagRequired("cluster-id")
	snowballCmd.AddCommand(snowball_describeClusterCmd)
}
