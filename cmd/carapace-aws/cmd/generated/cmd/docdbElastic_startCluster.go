package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_startClusterCmd = &cobra.Command{
	Use:   "start-cluster",
	Short: "Restarts the stopped elastic cluster that is specified by `clusterARN`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_startClusterCmd).Standalone()

	docdbElastic_startClusterCmd.Flags().String("cluster-arn", "", "The ARN identifier of the elastic cluster.")
	docdbElastic_startClusterCmd.MarkFlagRequired("cluster-arn")
	docdbElasticCmd.AddCommand(docdbElastic_startClusterCmd)
}
