package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_getHostCmd = &cobra.Command{
	Use:   "get-host",
	Short: "Returns the host ARN and details such as status, provider type, endpoint, and, if applicable, the VPC configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_getHostCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeconnections_getHostCmd).Standalone()

		codeconnections_getHostCmd.Flags().String("host-arn", "", "The Amazon Resource Name (ARN) of the requested host.")
		codeconnections_getHostCmd.MarkFlagRequired("host-arn")
	})
	codeconnectionsCmd.AddCommand(codeconnections_getHostCmd)
}
