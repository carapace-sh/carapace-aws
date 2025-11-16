package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_batchCheckLayerAvailabilityCmd = &cobra.Command{
	Use:   "batch-check-layer-availability",
	Short: "Checks the availability of one or more image layers that are within a repository in a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_batchCheckLayerAvailabilityCmd).Standalone()

	ecrPublic_batchCheckLayerAvailabilityCmd.Flags().String("layer-digests", "", "The digests of the image layers to check.")
	ecrPublic_batchCheckLayerAvailabilityCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID, or registry alias, associated with the public registry that contains the image layers to check.")
	ecrPublic_batchCheckLayerAvailabilityCmd.Flags().String("repository-name", "", "The name of the repository that's associated with the image layers to check.")
	ecrPublic_batchCheckLayerAvailabilityCmd.MarkFlagRequired("layer-digests")
	ecrPublic_batchCheckLayerAvailabilityCmd.MarkFlagRequired("repository-name")
	ecrPublicCmd.AddCommand(ecrPublic_batchCheckLayerAvailabilityCmd)
}
