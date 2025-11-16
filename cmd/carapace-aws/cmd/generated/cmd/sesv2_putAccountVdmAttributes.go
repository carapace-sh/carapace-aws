package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sesv2_putAccountVdmAttributesCmd = &cobra.Command{
	Use:   "put-account-vdm-attributes",
	Short: "Update your Amazon SES account VDM attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sesv2_putAccountVdmAttributesCmd).Standalone()

	sesv2_putAccountVdmAttributesCmd.Flags().String("vdm-attributes", "", "The VDM attributes that you wish to apply to your Amazon SES account.")
	sesv2_putAccountVdmAttributesCmd.MarkFlagRequired("vdm-attributes")
	sesv2Cmd.AddCommand(sesv2_putAccountVdmAttributesCmd)
}
