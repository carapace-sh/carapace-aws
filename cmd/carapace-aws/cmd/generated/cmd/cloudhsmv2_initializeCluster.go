package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudhsmv2_initializeClusterCmd = &cobra.Command{
	Use:   "initialize-cluster",
	Short: "Claims an CloudHSM cluster by submitting the cluster certificate issued by your issuing certificate authority (CA) and the CA's root certificate.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudhsmv2_initializeClusterCmd).Standalone()

	cloudhsmv2_initializeClusterCmd.Flags().String("cluster-id", "", "The identifier (ID) of the cluster that you are claiming.")
	cloudhsmv2_initializeClusterCmd.Flags().String("signed-cert", "", "The cluster certificate issued (signed) by your issuing certificate authority (CA).")
	cloudhsmv2_initializeClusterCmd.Flags().String("trust-anchor", "", "The issuing certificate of the issuing certificate authority (CA) that issued (signed) the cluster certificate.")
	cloudhsmv2_initializeClusterCmd.MarkFlagRequired("cluster-id")
	cloudhsmv2_initializeClusterCmd.MarkFlagRequired("signed-cert")
	cloudhsmv2_initializeClusterCmd.MarkFlagRequired("trust-anchor")
	cloudhsmv2Cmd.AddCommand(cloudhsmv2_initializeClusterCmd)
}
