package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datasync_listAgentsCmd = &cobra.Command{
	Use:   "list-agents",
	Short: "Returns a list of DataSync agents that belong to an Amazon Web Services account in the Amazon Web Services Region specified in the request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datasync_listAgentsCmd).Standalone()

	datasync_listAgentsCmd.Flags().String("max-results", "", "Specifies the maximum number of DataSync agents to list in a response.")
	datasync_listAgentsCmd.Flags().String("next-token", "", "Specifies an opaque string that indicates the position to begin the next list of results in the response.")
	datasyncCmd.AddCommand(datasync_listAgentsCmd)
}
