package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_getClusterCmd = &cobra.Command{
	Use:   "get-cluster",
	Short: "Returns information about a specific elastic cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_getClusterCmd).Standalone()

	docdbElastic_getClusterCmd.Flags().String("cluster-arn", "", "The ARN identifier of the elastic cluster.")
	docdbElastic_getClusterCmd.MarkFlagRequired("cluster-arn")
	docdbElasticCmd.AddCommand(docdbElastic_getClusterCmd)
}
