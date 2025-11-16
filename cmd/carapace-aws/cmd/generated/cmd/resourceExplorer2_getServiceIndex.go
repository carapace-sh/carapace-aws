package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_getServiceIndexCmd = &cobra.Command{
	Use:   "get-service-index",
	Short: "Retrieves information about the Resource Explorer index in the current Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_getServiceIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_getServiceIndexCmd).Standalone()

	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_getServiceIndexCmd)
}
