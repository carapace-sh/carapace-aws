package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_deleteContainerFleetCmd = &cobra.Command{
	Use:   "delete-container-fleet",
	Short: "**This API works with the following fleet types:** Container\n\nDeletes all resources and information related to a container fleet and shuts down currently running fleet instances, including those in remote locations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_deleteContainerFleetCmd).Standalone()

	gamelift_deleteContainerFleetCmd.Flags().String("fleet-id", "", "A unique identifier for the container fleet to delete.")
	gamelift_deleteContainerFleetCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_deleteContainerFleetCmd)
}
