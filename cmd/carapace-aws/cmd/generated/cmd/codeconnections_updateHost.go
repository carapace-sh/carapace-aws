package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeconnections_updateHostCmd = &cobra.Command{
	Use:   "update-host",
	Short: "Updates a specified host with the provided configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeconnections_updateHostCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeconnections_updateHostCmd).Standalone()

		codeconnections_updateHostCmd.Flags().String("host-arn", "", "The Amazon Resource Name (ARN) of the host to be updated.")
		codeconnections_updateHostCmd.Flags().String("provider-endpoint", "", "The URL or endpoint of the host to be updated.")
		codeconnections_updateHostCmd.Flags().String("vpc-configuration", "", "The VPC configuration of the host to be updated.")
		codeconnections_updateHostCmd.MarkFlagRequired("host-arn")
	})
	codeconnectionsCmd.AddCommand(codeconnections_updateHostCmd)
}
