package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codestarConnections_getHostCmd = &cobra.Command{
	Use:   "get-host",
	Short: "Returns the host ARN and details such as status, provider type, endpoint, and, if applicable, the VPC configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codestarConnections_getHostCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codestarConnections_getHostCmd).Standalone()

		codestarConnections_getHostCmd.Flags().String("host-arn", "", "The Amazon Resource Name (ARN) of the requested host.")
		codestarConnections_getHostCmd.MarkFlagRequired("host-arn")
	})
	codestarConnectionsCmd.AddCommand(codestarConnections_getHostCmd)
}
