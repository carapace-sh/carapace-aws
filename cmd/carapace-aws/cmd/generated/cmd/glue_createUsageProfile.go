package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var glue_createUsageProfileCmd = &cobra.Command{
	Use:   "create-usage-profile",
	Short: "Creates an Glue usage profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(glue_createUsageProfileCmd).Standalone()

	glue_createUsageProfileCmd.Flags().String("configuration", "", "A `ProfileConfiguration` object specifying the job and session values for the profile.")
	glue_createUsageProfileCmd.Flags().String("description", "", "A description of the usage profile.")
	glue_createUsageProfileCmd.Flags().String("name", "", "The name of the usage profile.")
	glue_createUsageProfileCmd.Flags().String("tags", "", "A list of tags applied to the usage profile.")
	glue_createUsageProfileCmd.MarkFlagRequired("configuration")
	glue_createUsageProfileCmd.MarkFlagRequired("name")
	glueCmd.AddCommand(glue_createUsageProfileCmd)
}
