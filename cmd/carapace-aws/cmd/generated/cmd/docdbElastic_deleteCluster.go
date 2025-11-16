package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var docdbElastic_deleteClusterCmd = &cobra.Command{
	Use:   "delete-cluster",
	Short: "Delete an elastic cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docdbElastic_deleteClusterCmd).Standalone()

	docdbElastic_deleteClusterCmd.Flags().String("cluster-arn", "", "The ARN identifier of the elastic cluster that is to be deleted.")
	docdbElastic_deleteClusterCmd.MarkFlagRequired("cluster-arn")
	docdbElasticCmd.AddCommand(docdbElastic_deleteClusterCmd)
}
