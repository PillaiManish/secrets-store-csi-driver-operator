package main

import (
	"k8s.io/utils/clock"
	"os"

	"github.com/openshift/library-go/pkg/controller/controllercmd"
	"github.com/spf13/cobra"
	"k8s.io/component-base/cli"

	"github.com/openshift/secrets-store-csi-driver-operator/pkg/operator"
	"github.com/openshift/secrets-store-csi-driver-operator/pkg/version"
)

func main() {
	command := NewOperatorCommand()
	code := cli.Run(command)
	os.Exit(code)
}

func NewOperatorCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "secrets-store-csi-driver-operator",
		Short: "OpenShift Secrets Store CSI Driver Operator",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
			os.Exit(1)
		},
	}

	ctrlCmd := controllercmd.NewControllerCommandConfig(
		"secrets-store-csi-driver-operator",
		version.Get(),
		operator.RunOperator,
		clock.RealClock{},
	).NewCommand()
	ctrlCmd.Use = "start"
	ctrlCmd.Short = "Start the Secrets Store CSI Driver Operator"

	cmd.AddCommand(ctrlCmd)

	return cmd
}
