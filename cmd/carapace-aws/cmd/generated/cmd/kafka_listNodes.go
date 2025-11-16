package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kafka_listNodesCmd = &cobra.Command{
	Use:   "list-nodes",
	Short: "Returns a list of the broker nodes in the cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kafka_listNodesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kafka_listNodesCmd).Standalone()

		kafka_listNodesCmd.Flags().String("cluster-arn", "", "The Amazon Resource Name (ARN) that uniquely identifies the cluster.")
		kafka_listNodesCmd.Flags().String("max-results", "", "The maximum number of results to return in the response.")
		kafka_listNodesCmd.Flags().String("next-token", "", "The paginated results marker.")
		kafka_listNodesCmd.MarkFlagRequired("cluster-arn")
	})
	kafkaCmd.AddCommand(kafka_listNodesCmd)
}
