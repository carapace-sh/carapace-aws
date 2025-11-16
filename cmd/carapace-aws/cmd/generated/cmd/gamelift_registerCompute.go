package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_registerComputeCmd = &cobra.Command{
	Use:   "register-compute",
	Short: "**This API works with the following fleet types:** Anywhere, Container\n\nRegisters a compute resource in an Amazon GameLift Servers Anywhere fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_registerComputeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(gamelift_registerComputeCmd).Standalone()

		gamelift_registerComputeCmd.Flags().String("certificate-path", "", "The path to a TLS certificate on your compute resource.")
		gamelift_registerComputeCmd.Flags().String("compute-name", "", "A descriptive label for the compute resource.")
		gamelift_registerComputeCmd.Flags().String("dns-name", "", "The DNS name of the compute resource.")
		gamelift_registerComputeCmd.Flags().String("fleet-id", "", "A unique identifier for the fleet to register the compute to.")
		gamelift_registerComputeCmd.Flags().String("ip-address", "", "The IP address of the compute resource.")
		gamelift_registerComputeCmd.Flags().String("location", "", "The name of a custom location to associate with the compute resource being registered.")
		gamelift_registerComputeCmd.MarkFlagRequired("compute-name")
		gamelift_registerComputeCmd.MarkFlagRequired("fleet-id")
	})
	gameliftCmd.AddCommand(gamelift_registerComputeCmd)
}
