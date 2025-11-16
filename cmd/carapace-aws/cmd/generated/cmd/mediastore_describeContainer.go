package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastore_describeContainerCmd = &cobra.Command{
	Use:   "describe-container",
	Short: "Retrieves the properties of the requested container.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastore_describeContainerCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastore_describeContainerCmd).Standalone()

		mediastore_describeContainerCmd.Flags().String("container-name", "", "The name of the container to query.")
	})
	mediastoreCmd.AddCommand(mediastore_describeContainerCmd)
}
