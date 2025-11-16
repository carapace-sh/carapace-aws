package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_getDefaultViewCmd = &cobra.Command{
	Use:   "get-default-view",
	Short: "Retrieves the Amazon Resource Name (ARN) of the view that is the default for the Amazon Web Services Region in which you call this operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_getDefaultViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_getDefaultViewCmd).Standalone()

	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_getDefaultViewCmd)
}
