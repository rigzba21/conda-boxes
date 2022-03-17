package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    //"k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/tools/clientcmd/api"
    "k8s.io/cli-runtime/pkg/genericclioptions"
)


var (
    exampleUsage = `
    # create a conda environment as a volume, given an environment file
    %[1]s conda create my-environment.yaml

    # list kubectl-conda managed environments
    %[1]s conda --list

    # delete a volume managed by kubectl-conda
    %[1]s conda delete --environment my-environment
`
    errNoContext = fmt.Errorf("no k8s context set")
)
type CondaOptions struct {
    configFlags *genericclioptions.ConfigFlags
    resultingContext *api.Context
    rawConfig api.Config
    args []string
    genericclioptions.IOStreams
}

func NewCondaOptions(streams genericclioptions.IOStreams) *CondaOptions {
    return &CondaOptions{
        configFlags: genericclioptions.NewConfigFlags(true),
        IOStreams: streams,
    }
}

func NewCmdConda(streams genericclioptions.IOStreams) *cobra.Command {
    options := NewCondaOptions(streams)

    cmd := &cobra.Command{
        Use: "conda [environment-file] [flags]",
        Short: "Create Conda Enviroment as Volume",
        Example: fmt.Sprintf(exampleUsage, "kubectl"),
        SilenceUsage: true,
        RunE: func(c *cobra.Command, args []string) error {
            if err := options.Run(c, args); err != nil {
                return err
            }
            
            return nil
        },
    }

    options.configFlags.AddFlags(cmd.Flags())
    return cmd
}

func (options *CondaOptions) Run(cmd *cobra.Command, args []string) error {
    options.args = args

    var err error
    options.rawConfig, err = options.configFlags.ToRawKubeConfigLoader().RawConfig()
    if err != nil {
        return err
    }

    fmt.Println("Running kubectl-conda...")

    _, exists := options.rawConfig.Contexts[options.rawConfig.CurrentContext]
    if !exists {
        return errNoContext
    }

    return nil
}
