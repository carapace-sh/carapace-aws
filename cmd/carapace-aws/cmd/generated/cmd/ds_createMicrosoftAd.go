package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_createMicrosoftAdCmd = &cobra.Command{
	Use:   "create-microsoft-ad",
	Short: "Creates a Microsoft AD directory in the Amazon Web Services Cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_createMicrosoftAdCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_createMicrosoftAdCmd).Standalone()

		ds_createMicrosoftAdCmd.Flags().String("description", "", "A description for the directory.")
		ds_createMicrosoftAdCmd.Flags().String("edition", "", "Managed Microsoft AD is available in two editions: `Standard` and `Enterprise`.")
		ds_createMicrosoftAdCmd.Flags().String("name", "", "The fully qualified domain name for the Managed Microsoft AD directory, such as `corp.example.com`.")
		ds_createMicrosoftAdCmd.Flags().String("network-type", "", "The network type for your domain.")
		ds_createMicrosoftAdCmd.Flags().String("password", "", "The password for the default administrative user named `Admin`.")
		ds_createMicrosoftAdCmd.Flags().String("short-name", "", "The NetBIOS name for your domain, such as `CORP`.")
		ds_createMicrosoftAdCmd.Flags().String("tags", "", "The tags to be assigned to the Managed Microsoft AD directory.")
		ds_createMicrosoftAdCmd.Flags().String("vpc-settings", "", "Contains VPC information for the [CreateDirectory]() or [CreateMicrosoftAD]() operation.")
		ds_createMicrosoftAdCmd.MarkFlagRequired("name")
		ds_createMicrosoftAdCmd.MarkFlagRequired("password")
		ds_createMicrosoftAdCmd.MarkFlagRequired("vpc-settings")
	})
	dsCmd.AddCommand(ds_createMicrosoftAdCmd)
}
