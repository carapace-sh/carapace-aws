package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateSdiSourceCmd = &cobra.Command{
	Use:   "update-sdi-source",
	Short: "Change some of the settings in an SdiSource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateSdiSourceCmd).Standalone()

	medialive_updateSdiSourceCmd.Flags().String("mode", "", "Include this parameter only if you want to change the name of the SdiSource.")
	medialive_updateSdiSourceCmd.Flags().String("name", "", "Include this parameter only if you want to change the name of the SdiSource.")
	medialive_updateSdiSourceCmd.Flags().String("sdi-source-id", "", "The ID of the SdiSource")
	medialive_updateSdiSourceCmd.Flags().String("type", "", "Include this parameter only if you want to change the mode.")
	medialive_updateSdiSourceCmd.MarkFlagRequired("sdi-source-id")
	medialiveCmd.AddCommand(medialive_updateSdiSourceCmd)
}
