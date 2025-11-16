package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_createConnectionCmd = &cobra.Command{
	Use:   "create-connection",
	Short: "Creates a connection that can then be given to other Amazon Web Services services like CodePipeline so that it can access third-party code repositories.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_createConnectionCmd).Standalone()

	codestarConnections_createConnectionCmd.Flags().String("connection-name", "", "The name of the connection to be created.")
	codestarConnections_createConnectionCmd.Flags().String("host-arn", "", "The Amazon Resource Name (ARN) of the host associated with the connection to be created.")
	codestarConnections_createConnectionCmd.Flags().String("provider-type", "", "The name of the external provider where your third-party code repository is configured.")
	codestarConnections_createConnectionCmd.Flags().String("tags", "", "The key-value pair to use when tagging the resource.")
	codestarConnections_createConnectionCmd.MarkFlagRequired("connection-name")
	codestarConnectionsCmd.AddCommand(codestarConnections_createConnectionCmd)
}
