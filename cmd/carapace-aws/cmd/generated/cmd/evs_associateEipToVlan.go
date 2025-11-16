package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evs_associateEipToVlanCmd = &cobra.Command{
	Use:   "associate-eip-to-vlan",
	Short: "Associates an Elastic IP address with a public HCX VLAN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evs_associateEipToVlanCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evs_associateEipToVlanCmd).Standalone()

		evs_associateEipToVlanCmd.Flags().String("allocation-id", "", "The Elastic IP address allocation ID.")
		evs_associateEipToVlanCmd.Flags().String("client-token", "", "This parameter is not used in Amazon EVS currently.")
		evs_associateEipToVlanCmd.Flags().String("environment-id", "", "A unique ID for the environment containing the VLAN that the Elastic IP address associates with.")
		evs_associateEipToVlanCmd.Flags().String("vlan-name", "", "The name of the VLAN.")
		evs_associateEipToVlanCmd.MarkFlagRequired("allocation-id")
		evs_associateEipToVlanCmd.MarkFlagRequired("environment-id")
		evs_associateEipToVlanCmd.MarkFlagRequired("vlan-name")
	})
	evsCmd.AddCommand(evs_associateEipToVlanCmd)
}
