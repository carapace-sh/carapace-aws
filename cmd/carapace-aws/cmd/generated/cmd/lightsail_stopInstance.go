package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_stopInstanceCmd = &cobra.Command{
	Use:   "stop-instance",
	Short: "Stops a specific Amazon Lightsail instance that is currently running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_stopInstanceCmd).Standalone()

	lightsail_stopInstanceCmd.Flags().String("force", "", "When set to `True`, forces a Lightsail instance that is stuck in a `stopping` state to stop.")
	lightsail_stopInstanceCmd.Flags().String("instance-name", "", "The name of the instance (a virtual private server) to stop.")
	lightsail_stopInstanceCmd.MarkFlagRequired("instance-name")
	lightsailCmd.AddCommand(lightsail_stopInstanceCmd)
}
