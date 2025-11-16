package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listClientVpcConnectionsCmd = &cobra.Command{
	Use:   "list-client-vpc-connections",
	Short: "Returns a list of all the VPC connections in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listClientVpcConnectionsCmd).Standalone()

	kafka_listClientVpcConnectionsCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) of the cluster.")
	kafka_listClientVpcConnectionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	kafka_listClientVpcConnectionsCmd.Flags().String("next-token", "", "The paginated results marker.")
	kafka_listClientVpcConnectionsCmd.MarkFlagRequired("cluster-arn")
	kafkaCmd.AddCommand(kafka_listClientVpcConnectionsCmd)
}
