package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createNodeRegistrationScriptCmd = &cobra.Command{
	Use:   "create-node-registration-script",
	Short: "Create the Register Node script for all the nodes intended for a specific Cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createNodeRegistrationScriptCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_createNodeRegistrationScriptCmd).Standalone()

		medialive_createNodeRegistrationScriptCmd.Flags().String("cluster-id", "", "The ID of the cluster")
		medialive_createNodeRegistrationScriptCmd.Flags().String("id", "", "If you're generating a re-registration script for an already existing node, this is where you provide the id.")
		medialive_createNodeRegistrationScriptCmd.Flags().String("name", "", "Specify a pattern for MediaLive Anywhere to use to assign a name to each Node in the Cluster.")
		medialive_createNodeRegistrationScriptCmd.Flags().String("node-interface-mappings", "", "Documentation update needed")
		medialive_createNodeRegistrationScriptCmd.Flags().String("request-id", "", "An ID that you assign to a create request.")
		medialive_createNodeRegistrationScriptCmd.Flags().String("role", "", "The initial role of the Node in the Cluster.")
		medialive_createNodeRegistrationScriptCmd.MarkFlagRequired("cluster-id")
	})
	medialiveCmd.AddCommand(medialive_createNodeRegistrationScriptCmd)
}
