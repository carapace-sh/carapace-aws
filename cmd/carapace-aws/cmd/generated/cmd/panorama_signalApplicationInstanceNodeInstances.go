package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var panorama_signalApplicationInstanceNodeInstancesCmd = &cobra.Command{
	Use:   "signal-application-instance-node-instances",
	Short: "Signal camera nodes to stop or resume.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(panorama_signalApplicationInstanceNodeInstancesCmd).Standalone()

	panorama_signalApplicationInstanceNodeInstancesCmd.Flags().String("application-instance-id", "", "An application instance ID.")
	panorama_signalApplicationInstanceNodeInstancesCmd.Flags().String("node-signals", "", "A list of signals.")
	panorama_signalApplicationInstanceNodeInstancesCmd.MarkFlagRequired("application-instance-id")
	panorama_signalApplicationInstanceNodeInstancesCmd.MarkFlagRequired("node-signals")
	panoramaCmd.AddCommand(panorama_signalApplicationInstanceNodeInstancesCmd)
}
