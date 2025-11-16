package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_listInstancesCmd = &cobra.Command{
	Use:   "list-instances",
	Short: "This API is in preview release for Amazon Connect and is subject to change.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_listInstancesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_listInstancesCmd).Standalone()

		connect_listInstancesCmd.Flags().String("max-results", "", "The maximum number of results to return per page.")
		connect_listInstancesCmd.Flags().String("next-token", "", "The token for the next set of results.")
	})
	connectCmd.AddCommand(connect_listInstancesCmd)
}
