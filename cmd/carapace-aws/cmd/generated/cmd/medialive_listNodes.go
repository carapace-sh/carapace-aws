package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listNodesCmd = &cobra.Command{
	Use:   "list-nodes",
	Short: "Retrieve the list of Nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listNodesCmd).Standalone()

	medialive_listNodesCmd.Flags().String("cluster-id", "", "The ID of the cluster")
	medialive_listNodesCmd.Flags().String("max-results", "", "The maximum number of items to return.")
	medialive_listNodesCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	medialive_listNodesCmd.MarkFlagRequired("cluster-id")
	medialiveCmd.AddCommand(medialive_listNodesCmd)
}
