package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_disassociateEipFromVlanCmd = &cobra.Command{
	Use:   "disassociate-eip-from-vlan",
	Short: "Disassociates an Elastic IP address from a public HCX VLAN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_disassociateEipFromVlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evs_disassociateEipFromVlanCmd).Standalone()

		evs_disassociateEipFromVlanCmd.Flags().String("association-id", "", "A unique ID for the Elastic IP address association.")
		evs_disassociateEipFromVlanCmd.Flags().String("client-token", "", "This parameter is not used in Amazon EVS currently.")
		evs_disassociateEipFromVlanCmd.Flags().String("environment-id", "", "A unique ID for the environment containing the VLAN that the Elastic IP address disassociates from.")
		evs_disassociateEipFromVlanCmd.Flags().String("vlan-name", "", "The name of the VLAN.")
		evs_disassociateEipFromVlanCmd.MarkFlagRequired("association-id")
		evs_disassociateEipFromVlanCmd.MarkFlagRequired("environment-id")
		evs_disassociateEipFromVlanCmd.MarkFlagRequired("vlan-name")
	})
	evsCmd.AddCommand(evs_disassociateEipFromVlanCmd)
}
