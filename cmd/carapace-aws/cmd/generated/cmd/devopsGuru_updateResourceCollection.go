package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devopsGuru_updateResourceCollectionCmd = &cobra.Command{
	Use:   "update-resource-collection",
	Short: "Updates the collection of resources that DevOps Guru analyzes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devopsGuru_updateResourceCollectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devopsGuru_updateResourceCollectionCmd).Standalone()

		devopsGuru_updateResourceCollectionCmd.Flags().String("action", "", "Specifies if the resource collection in the request is added or deleted to the resource collection.")
		devopsGuru_updateResourceCollectionCmd.Flags().String("resource-collection", "", "")
		devopsGuru_updateResourceCollectionCmd.MarkFlagRequired("action")
		devopsGuru_updateResourceCollectionCmd.MarkFlagRequired("resource-collection")
	})
	devopsGuruCmd.AddCommand(devopsGuru_updateResourceCollectionCmd)
}
