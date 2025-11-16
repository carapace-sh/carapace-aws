package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_listNodesCmd = &cobra.Command{
	Use:   "list-nodes",
	Short: "Takes in filters and returns a list of managed nodes matching the filter criteria.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_listNodesCmd).Standalone()

	ssm_listNodesCmd.Flags().String("filters", "", "One or more filters.")
	ssm_listNodesCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_listNodesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_listNodesCmd.Flags().String("sync-name", "", "The name of the Amazon Web Services managed resource data sync to retrieve information about.")
	ssmCmd.AddCommand(ssm_listNodesCmd)
}
