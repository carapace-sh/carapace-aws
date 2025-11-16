package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createSdiSourceCmd = &cobra.Command{
	Use:   "create-sdi-source",
	Short: "Create an SdiSource for each video source that uses the SDI protocol.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createSdiSourceCmd).Standalone()

	medialive_createSdiSourceCmd.Flags().String("mode", "", "Applies only if the type is QUAD.")
	medialive_createSdiSourceCmd.Flags().String("name", "", "Specify a name that is unique in the AWS account.")
	medialive_createSdiSourceCmd.Flags().String("request-id", "", "An ID that you assign to a create request.")
	medialive_createSdiSourceCmd.Flags().String("tags", "", "A collection of key-value pairs.")
	medialive_createSdiSourceCmd.Flags().String("type", "", "Specify the type of the SDI source: SINGLE: The source is a single-link source.")
	medialiveCmd.AddCommand(medialive_createSdiSourceCmd)
}
