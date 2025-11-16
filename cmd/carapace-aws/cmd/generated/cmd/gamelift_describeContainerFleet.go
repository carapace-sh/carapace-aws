package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var gamelift_describeContainerFleetCmd = &cobra.Command{
	Use:   "describe-container-fleet",
	Short: "**This API works with the following fleet types:** Container\n\nRetrieves the properties for a container fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gamelift_describeContainerFleetCmd).Standalone()

	gamelift_describeContainerFleetCmd.Flags().String("fleet-id", "", "A unique identifier for the container fleet to retrieve.")
	gamelift_describeContainerFleetCmd.MarkFlagRequired("fleet-id")
	gameliftCmd.AddCommand(gamelift_describeContainerFleetCmd)
}
