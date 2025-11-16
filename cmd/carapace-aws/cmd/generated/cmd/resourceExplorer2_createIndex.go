package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourceExplorer2_createIndexCmd = &cobra.Command{
	Use:   "create-index",
	Short: "Turns on Amazon Web Services Resource Explorer in the Amazon Web Services Region in which you called this operation by creating an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourceExplorer2_createIndexCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resourceExplorer2_createIndexCmd).Standalone()

		resourceExplorer2_createIndexCmd.Flags().String("client-token", "", "This value helps ensure idempotency.")
		resourceExplorer2_createIndexCmd.Flags().String("tags", "", "The specified tags are attached only to the index created in this Amazon Web Services Region.")
	})
	resourceExplorer2Cmd.AddCommand(resourceExplorer2_createIndexCmd)
}
