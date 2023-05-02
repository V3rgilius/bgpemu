package lab

import (
	"fmt"
	"os"
	"strings"

	spb "github.com/p3rdy/bgpemu/proto/scene"
	"google.golang.org/protobuf/encoding/prototext"
	"sigs.k8s.io/yaml"
)

func LoadScene(path string) (*spb.Scene, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	pds := &spb.Scene{}
	switch {
	case strings.HasSuffix(path, ".yaml"), strings.HasSuffix(path, ".yml"):
		jsonBytes, err := yaml.YAMLToJSON(b)
		if err != nil {
			return nil, fmt.Errorf("could not parse yaml: %v", err)
		}
		if err := protojsonUnmarshaller.Unmarshal(jsonBytes, pds); err != nil {
			return nil, fmt.Errorf("could not parse json: %v", err)
		}
	default:
		if err := prototext.Unmarshal(b, pds); err != nil {
			return nil, err
		}
	}
	return pds, nil
}

func Behave(scene *spb.Scene) error {
	return nil
}

func execSteps(steps []*spb.Step) error {
	return nil
}

func execStep(step *spb.Step) error {
	return nil
}

func execCmds(cmds *spb.Commands) error {
	return nil
}

func transportFile(file *spb.FileTrans) error {
	return nil
}
