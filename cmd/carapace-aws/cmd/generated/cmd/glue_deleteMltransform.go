package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_deleteMltransformCmd = &cobra.Command{
	Use:   "delete-mltransform",
	Short: "Deletes an Glue machine learning transform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_deleteMltransformCmd).Standalone()

	glue_deleteMltransformCmd.Flags().String("transform-id", "", "The unique identifier of the transform to delete.")
	glue_deleteMltransformCmd.MarkFlagRequired("transform-id")
	glueCmd.AddCommand(glue_deleteMltransformCmd)
}
