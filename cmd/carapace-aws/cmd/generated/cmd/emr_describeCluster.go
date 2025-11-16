package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_describeClusterCmd = &cobra.Command{
	Use:   "describe-cluster",
	Short: "Provides cluster-level details including status, hardware and software configuration, VPC settings, and so on.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_describeClusterCmd).Standalone()

	emr_describeClusterCmd.Flags().String("cluster-id", "", "The identifier of the cluster to describe.")
	emr_describeClusterCmd.MarkFlagRequired("cluster-id")
	emrCmd.AddCommand(emr_describeClusterCmd)
}
