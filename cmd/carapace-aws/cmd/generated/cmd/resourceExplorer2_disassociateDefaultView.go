package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_disassociateDefaultViewCmd = &cobra.Command{
	Use:   "disassociate-default-view",
	Short: "After you call this operation, the affected Amazon Web Services Region no longer has a default view.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_disassociateDefaultViewCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_disassociateDefaultViewCmd).Standalone()

	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_disassociateDefaultViewCmd)
}
