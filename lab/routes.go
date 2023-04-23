package lab

import (
	"context"
	"fmt"
	"os"
	"strings"

	rtpb "github.com/p3rdy/bgpemu/proto/routes"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"
)

var protojsonUnmarshaller = protojson.UnmarshalOptions{
	AllowPartial:   true,
	DiscardUnknown: false,
}

// LoadRoutes loads a Topology from path and parse all subtopos.
func LoadRoutes(path string) (*rtpb.RouteDeployment, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	rd := &rtpb.RouteDeployment{}
	switch {
	case strings.HasSuffix(path, ".yaml"), strings.HasSuffix(path, ".yml"):
		jsonBytes, err := yaml.YAMLToJSON(b)
		if err != nil {
			return nil, fmt.Errorf("could not parse yaml: %v", err)
		}
		if err := protojsonUnmarshaller.Unmarshal(jsonBytes, rd); err != nil {
			return nil, fmt.Errorf("could not parse json: %v", err)
		}
	default:
		if err := prototext.Unmarshal(b, rd); err != nil {
			return nil, err
		}
	}
	return rd, nil
}

func DeployRoutes(rd *rtpb.RouteDeployment, m *Manager) error {
	// 获取路由文件中设备对应的Pod的gRPC接口
	// 创建gRPC连接
	// 构造，调用
	for _, r := range rd.Routes {
		svcName := fmt.Sprintf("service-%s", r.Name)
		service, err := m.kClient.CoreV1().Services(rd.TopoName).Get(context.Background(), svcName, metav1.GetOptions{})
		if err != nil {
			return err
		}
		if(service.Spec.Ports)
	}
	return nil
}
