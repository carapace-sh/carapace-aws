package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_deleteServiceCmd = &cobra.Command{
	Use:   "delete-service",
	Short: "Deletes a specified service and all associated service attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_deleteServiceCmd).Standalone()

	servicediscovery_deleteServiceCmd.Flags().String("id", "", "The ID or Amazon Resource Name (ARN) of the service that you want to delete.")
	servicediscovery_deleteServiceCmd.MarkFlagRequired("id")
	servicediscoveryCmd.AddCommand(servicediscovery_deleteServiceCmd)
}
