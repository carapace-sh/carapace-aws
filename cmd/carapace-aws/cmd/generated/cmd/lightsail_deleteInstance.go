package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_deleteInstanceCmd = &cobra.Command{
	Use:   "delete-instance",
	Short: "Deletes an Amazon Lightsail instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_deleteInstanceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_deleteInstanceCmd).Standalone()

		lightsail_deleteInstanceCmd.Flags().String("force-delete-add-ons", "", "A Boolean value to indicate whether to delete all add-ons for the instance.")
		lightsail_deleteInstanceCmd.Flags().String("instance-name", "", "The name of the instance to delete.")
		lightsail_deleteInstanceCmd.MarkFlagRequired("instance-name")
	})
	lightsailCmd.AddCommand(lightsail_deleteInstanceCmd)
}
