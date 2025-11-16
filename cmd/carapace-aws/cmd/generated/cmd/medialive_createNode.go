package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createNodeCmd = &cobra.Command{
	Use:   "create-node",
	Short: "Create a Node in the specified Cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createNodeCmd).Standalone()

	medialive_createNodeCmd.Flags().String("cluster-id", "", "The ID of the cluster.")
	medialive_createNodeCmd.Flags().String("name", "", "The user-specified name of the Node to be created.")
	medialive_createNodeCmd.Flags().String("node-interface-mappings", "", "Documentation update needed")
	medialive_createNodeCmd.Flags().String("request-id", "", "An ID that you assign to a create request.")
	medialive_createNodeCmd.Flags().String("role", "", "The initial role of the Node in the Cluster.")
	medialive_createNodeCmd.Flags().String("tags", "", "A collection of key-value pairs.")
	medialive_createNodeCmd.MarkFlagRequired("cluster-id")
	medialiveCmd.AddCommand(medialive_createNodeCmd)
}
