package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_getInstanceStateCmd = &cobra.Command{
	Use:   "get-instance-state",
	Short: "Returns the state of a specific instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_getInstanceStateCmd).Standalone()

	lightsail_getInstanceStateCmd.Flags().String("instance-name", "", "The name of the instance to get state information about.")
	lightsail_getInstanceStateCmd.MarkFlagRequired("instance-name")
	lightsailCmd.AddCommand(lightsail_getInstanceStateCmd)
}
