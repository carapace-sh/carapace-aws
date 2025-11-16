package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_startBlueprintRunCmd = &cobra.Command{
	Use:   "start-blueprint-run",
	Short: "Starts a new run of the specified blueprint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_startBlueprintRunCmd).Standalone()

	glue_startBlueprintRunCmd.Flags().String("blueprint-name", "", "The name of the blueprint.")
	glue_startBlueprintRunCmd.Flags().String("parameters", "", "Specifies the parameters as a `BlueprintParameters` object.")
	glue_startBlueprintRunCmd.Flags().String("role-arn", "", "Specifies the IAM role used to create the workflow.")
	glue_startBlueprintRunCmd.MarkFlagRequired("blueprint-name")
	glue_startBlueprintRunCmd.MarkFlagRequired("role-arn")
	glueCmd.AddCommand(glue_startBlueprintRunCmd)
}
