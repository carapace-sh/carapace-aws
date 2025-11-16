package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_getIndexCmd = &cobra.Command{
	Use:   "get-index",
	Short: "Retrieves details about the Amazon Web Services Resource Explorer index in the Amazon Web Services Region in which you invoked the operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_getIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_getIndexCmd).Standalone()

	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_getIndexCmd)
}
