package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managedblockchain_listNodesCmd = &cobra.Command{
	Use:   "list-nodes",
	Short: "Returns information about the nodes within a network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managedblockchain_listNodesCmd).Standalone()

	managedblockchain_listNodesCmd.Flags().String("max-results", "", "The maximum number of nodes to list.")
	managedblockchain_listNodesCmd.Flags().String("member-id", "", "The unique identifier of the member who owns the nodes to list.")
	managedblockchain_listNodesCmd.Flags().String("network-id", "", "The unique identifier of the network for which to list nodes.")
	managedblockchain_listNodesCmd.Flags().String("next-token", "", "The pagination token that indicates the next set of results to retrieve.")
	managedblockchain_listNodesCmd.Flags().String("status", "", "An optional status specifier.")
	managedblockchain_listNodesCmd.MarkFlagRequired("network-id")
	managedblockchainCmd.AddCommand(managedblockchain_listNodesCmd)
}
