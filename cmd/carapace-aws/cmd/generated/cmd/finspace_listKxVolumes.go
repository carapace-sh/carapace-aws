package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspace_listKxVolumesCmd = &cobra.Command{
	Use:   "list-kx-volumes",
	Short: "Lists all the volumes in a kdb environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspace_listKxVolumesCmd).Standalone()

	finspace_listKxVolumesCmd.Flags().String("environment-id", "", "A unique identifier for the kdb environment, whose clusters can attach to the volume.")
	finspace_listKxVolumesCmd.Flags().String("max-results", "", "The maximum number of results to return in this request.")
	finspace_listKxVolumesCmd.Flags().String("next-token", "", "A token that indicates where a results page should begin.")
	finspace_listKxVolumesCmd.Flags().String("volume-type", "", "The type of file system volume.")
	finspace_listKxVolumesCmd.MarkFlagRequired("environment-id")
	finspaceCmd.AddCommand(finspace_listKxVolumesCmd)
}
