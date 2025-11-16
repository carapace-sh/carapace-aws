package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_stopClusterCmd = &cobra.Command{
	Use:   "stop-cluster",
	Short: "Stops the running elastic cluster that is specified by `clusterArn`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_stopClusterCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(docdbElastic_stopClusterCmd).Standalone()

		docdbElastic_stopClusterCmd.Flags().String("cluster-arn", "", "The ARN identifier of the elastic cluster.")
		docdbElastic_stopClusterCmd.MarkFlagRequired("cluster-arn")
	})
	docdbElasticCmd.AddCommand(docdbElastic_stopClusterCmd)
}
