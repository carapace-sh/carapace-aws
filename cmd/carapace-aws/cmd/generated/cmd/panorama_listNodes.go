package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_listNodesCmd = &cobra.Command{
	Use:   "list-nodes",
	Short: "Returns a list of nodes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_listNodesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(panorama_listNodesCmd).Standalone()

		panorama_listNodesCmd.Flags().String("category", "", "Search for nodes by category.")
		panorama_listNodesCmd.Flags().String("max-results", "", "The maximum number of nodes to return in one page of results.")
		panorama_listNodesCmd.Flags().String("next-token", "", "Specify the pagination token from a previous request to retrieve the next page of results.")
		panorama_listNodesCmd.Flags().String("owner-account", "", "Search for nodes by the account ID of the nodes' owner.")
		panorama_listNodesCmd.Flags().String("package-name", "", "Search for nodes by name.")
		panorama_listNodesCmd.Flags().String("package-version", "", "Search for nodes by version.")
		panorama_listNodesCmd.Flags().String("patch-version", "", "Search for nodes by patch version.")
	})
	panoramaCmd.AddCommand(panorama_listNodesCmd)
}
