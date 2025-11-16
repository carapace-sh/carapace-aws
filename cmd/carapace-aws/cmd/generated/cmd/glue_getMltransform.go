package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_getMltransformCmd = &cobra.Command{
	Use:   "get-mltransform",
	Short: "Gets an Glue machine learning transform artifact and all its corresponding metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_getMltransformCmd).Standalone()

	glue_getMltransformCmd.Flags().String("transform-id", "", "The unique identifier of the transform, generated at the time that the transform was created.")
	glue_getMltransformCmd.MarkFlagRequired("transform-id")
	glueCmd.AddCommand(glue_getMltransformCmd)
}
