package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var schemas_createDiscovererCmd = &cobra.Command{
	Use:   "create-discoverer",
	Short: "Creates a discoverer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(schemas_createDiscovererCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(schemas_createDiscovererCmd).Standalone()

		schemas_createDiscovererCmd.Flags().String("cross-account", "", "Support discovery of schemas in events sent to the bus from another account.")
		schemas_createDiscovererCmd.Flags().String("description", "", "A description for the discoverer.")
		schemas_createDiscovererCmd.Flags().String("source-arn", "", "The ARN of the event bus.")
		schemas_createDiscovererCmd.Flags().String("tags", "", "Tags associated with the resource.")
		schemas_createDiscovererCmd.MarkFlagRequired("source-arn")
	})
	schemasCmd.AddCommand(schemas_createDiscovererCmd)
}
