package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var storagegateway_describeSmbsettingsCmd = &cobra.Command{
	Use:   "describe-smbsettings",
	Short: "Gets a description of a Server Message Block (SMB) file share settings from a file gateway.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(storagegateway_describeSmbsettingsCmd).Standalone()

	storagegateway_describeSmbsettingsCmd.Flags().String("gateway-arn", "", "")
	storagegateway_describeSmbsettingsCmd.MarkFlagRequired("gateway-arn")
	storagegatewayCmd.AddCommand(storagegateway_describeSmbsettingsCmd)
}
