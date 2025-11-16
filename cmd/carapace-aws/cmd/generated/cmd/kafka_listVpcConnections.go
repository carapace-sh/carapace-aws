package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listVpcConnectionsCmd = &cobra.Command{
	Use:   "list-vpc-connections",
	Short: "Returns a list of all the VPC connections in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listVpcConnectionsCmd).Standalone()

	kafka_listVpcConnectionsCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
	kafka_listVpcConnectionsCmd.Flags().String("next-token", "", "The paginated results marker.")
	kafkaCmd.AddCommand(kafka_listVpcConnectionsCmd)
}
