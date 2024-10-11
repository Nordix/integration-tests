// Code generated by gotestmd DO NOT EDIT.
package features_ovs

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/ovs"
)

type Suite struct {
	base.Suite
	ovsSuite ovs.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.ovsSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
}
func (s *Suite) TestAnnotated_namespace() {
	r := s.Runner("../deployments-k8s/examples/features/annotated-namespace")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-annotated-namespace`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/annotated-namespace?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-annotated-namespace`)
	r.Run(`kubectl annotate ns ns-annotated-namespace networkservicemesh.io=kernel://annotated-namespace/nsm-1`)
	r.Run(`kubectl apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/c1632838c22fdf074abe49ac7f5177eb5937c02b/examples/features/annotated-namespace/client.yaml`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-annotated-namespace`)
	r.Run(`kubectl exec deployments/alpine -n ns-annotated-namespace -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-annotated-namespace -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestChange_nse_dynamically() {
	r := s.Runner("../deployments-k8s/examples/features/change-nse-dynamically")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-change-nse-dynamically`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/change-nse-dynamically?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/c1632838c22fdf074abe49ac7f5177eb5937c02b/examples/features/change-nse-dynamically/blue-netsvc.yaml`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-change-nse-dynamically`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=blue-nse -n ns-change-nse-dynamically`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=green-nse -n ns-change-nse-dynamically`)
	r.Run(`kubectl exec pods/alpine -n ns-change-nse-dynamically -- ping -c 4 172.16.2.100`)
	r.Run(`kubectl exec pods/blue-nse -n ns-change-nse-dynamically -- ping -c 4 172.16.2.101`)
	r.Run(`kubectl apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/c1632838c22fdf074abe49ac7f5177eb5937c02b/examples/features/change-nse-dynamically/green-netsvc.yaml`)
	r.Run(`kubectl exec pods/alpine -n ns-change-nse-dynamically -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec pods/green-nse -n ns-change-nse-dynamically -- ping -c 4 172.16.1.101`)
	r.Run(`kubectl apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/c1632838c22fdf074abe49ac7f5177eb5937c02b/examples/features/change-nse-dynamically/blue-netsvc.yaml`)
	r.Run(`kubectl exec pods/alpine -n ns-change-nse-dynamically -- ping -c 4 172.16.2.100`)
	r.Run(`kubectl exec pods/blue-nse -n ns-change-nse-dynamically -- ping -c 4 172.16.2.101`)
}
func (s *Suite) TestDns() {
	r := s.Runner("../deployments-k8s/examples/features/dns")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-dns`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/dns?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=5m pod dnsutils -n ns-dns`)
	r.Run(`kubectl wait --for=condition=ready --timeout=5m pod -l app=nse-kernel -n ns-dns`)
	r.Run(`kubectl exec pods/dnsutils -c dnsutils -n ns-dns -- nslookup -norec -nodef my.coredns.service`)
	r.Run(`kubectl exec pods/dnsutils -c dnsutils -n ns-dns -- ping -c 4 my.coredns.service`)
	r.Run(`kubectl exec pods/dnsutils -c dnsutils -n ns-dns -- dig kubernetes.default A kubernetes.default AAAA | grep "kubernetes.default.svc.cluster.local"`)
}
func (s *Suite) TestKernel2IP2Kernel_dual_stack() {
	r := s.Runner("../deployments-k8s/examples/features/dual-stack/Kernel2IP2Kernel_dual_stack")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2ip2kernel-dual-stack`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/dual-stack/Kernel2IP2Kernel_dual_stack?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2ip2kernel-dual-stack`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2ip2kernel-dual-stack`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2ip2kernel-dual-stack -- ping -c 4 2001:db8::`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-kernel2ip2kernel-dual-stack -- ping -c 4 2001:db8::1`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2ip2kernel-dual-stack -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-kernel2ip2kernel-dual-stack -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestKernel2Kernel_dual_stack() {
	r := s.Runner("../deployments-k8s/examples/features/dual-stack/Kernel2Kernel_dual_stack")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2kernel-dual-stack`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/dual-stack/Kernel2Kernel_dual_stack?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2kernel-dual-stack`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2kernel-dual-stack`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2kernel-dual-stack -- ping -c 4 2001:db8::`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-kernel2kernel-dual-stack -- ping -c 4 2001:db8::1`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2kernel-dual-stack -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-kernel2kernel-dual-stack -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestExclude_prefixes() {
	r := s.Runner("../deployments-k8s/examples/features/exclude-prefixes")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete configmap excluded-prefixes-config` + "\n" + `kubectl delete ns ns-exclude-prefixes`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/exclude-prefixes/configmap?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/exclude-prefixes?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-exclude-prefixes`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-exclude-prefixes`)
	r.Run(`kubectl exec pods/alpine -n ns-exclude-prefixes -- ping -c 4 172.16.1.200`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-exclude-prefixes -- ping -c 4 172.16.1.203`)
}
func (s *Suite) TestExclude_prefixes_client() {
	r := s.Runner("../deployments-k8s/examples/features/exclude-prefixes-client")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-exclude-prefixes-client`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/exclude-prefixes-client?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-exclude-prefixes-client`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel-1 -n ns-exclude-prefixes-client`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel-2 -n ns-exclude-prefixes-client`)
	r.Run(`kubectl exec pods/alpine -n ns-exclude-prefixes-client -- ping -c 4 172.16.1.96`)
	r.Run(`kubectl exec pods/alpine -n ns-exclude-prefixes-client -- ping -c 4 172.16.1.98`)
	r.Run(`kubectl exec deployments/nse-kernel-1 -n ns-exclude-prefixes-client -- ping -c 4 172.16.1.97`)
	r.Run(`kubectl exec deployments/nse-kernel-2 -n ns-exclude-prefixes-client -- ping -c 4 172.16.1.99`)
}
func (s *Suite) TestKernel2IP2Kernel_ipv6() {
	r := s.Runner("../deployments-k8s/examples/features/ipv6/Kernel2IP2Kernel_ipv6")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2ip2kernel-ipv6`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/ipv6/Kernel2IP2Kernel_ipv6?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2ip2kernel-ipv6`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2ip2kernel-ipv6`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2ip2kernel-ipv6 -- ping -c 4 2001:db8::`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-kernel2ip2kernel-ipv6 -- ping -c 4 2001:db8::1`)
}
func (s *Suite) TestKernel2Kernel_ipv6() {
	r := s.Runner("../deployments-k8s/examples/features/ipv6/Kernel2Kernel_ipv6")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2kernel-ipv6`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/ipv6/Kernel2Kernel_ipv6?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2kernel-ipv6`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2kernel-ipv6`)
	r.Run(`kubectl exec pods/alpine -n ns-kernel2kernel-ipv6 -- ping -c 4 2001:db8::`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-kernel2kernel-ipv6 -- ping -c 4 2001:db8::1`)
}
func (s *Suite) TestMultiple_services() {
	r := s.Runner("../deployments-k8s/examples/features/multiple-services")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-multiple-services`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/multiple-services?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-multiple-services`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel-1 -n ns-multiple-services`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel-2 -n ns-multiple-services`)
	r.Run(`kubectl exec pods/alpine -n ns-multiple-services -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec pods/nse-kernel-1 -n ns-multiple-services -- ping -c 4 172.16.1.101`)
	r.Run(`kubectl exec pods/alpine -n ns-multiple-services -- ping -c 4 172.16.2.100`)
	r.Run(`kubectl exec pods/nse-kernel-2 -n ns-multiple-services -- ping -c 4 172.16.2.101`)
}
func (s *Suite) TestMutually_aware_nses() {
	r := s.Runner("../deployments-k8s/examples/features/mutually-aware-nses")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-mutually-aware-nses`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/mutually-aware-nses?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-kernel -n ns-mutually-aware-nses`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel-1 -n ns-mutually-aware-nses`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel-2 -n ns-mutually-aware-nses`)
	r.Run(`kubectl exec deployments/nsc-kernel -n ns-mutually-aware-nses -- apk update` + "\n" + `kubectl exec deployments/nsc-kernel -n ns-mutually-aware-nses -- apk add iproute2`)
	r.Run(`result=$(kubectl exec deployments/nsc-kernel -n ns-mutually-aware-nses -- ip r get 172.16.1.100 from 172.16.1.101 ipproto tcp dport 6666)` + "\n" + `echo ${result}` + "\n" + `echo ${result} | grep -E -q "172.16.1.100 from 172.16.1.101 dev nsm-1"`)
	r.Run(`result=$(kubectl exec deployments/nsc-kernel -n ns-mutually-aware-nses -- ip r get 172.16.1.100 from 172.16.1.101 ipproto udp dport 5555)` + "\n" + `echo ${result}` + "\n" + `echo ${result} | grep -E -q "172.16.1.100 from 172.16.1.101 dev nsm-2"`)
}
func (s *Suite) TestOpa() {
	r := s.Runner("../deployments-k8s/examples/features/opa")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-opa`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/opa?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-kernel -n ns-opa`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-opa`)
	r.Run(`kubectl logs deployments/nsc-kernel -n ns-opa | grep "PermissionDenied desc = no sufficient privileges"`)
}
func (s *Suite) TestPolicy_based_routing() {
	r := s.Runner("../deployments-k8s/examples/features/policy-based-routing")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-policy-based-routing`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/policy-based-routing?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nettools -n ns-policy-based-routing`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-policy-based-routing`)
	r.Run(`kubectl exec pods/nettools -n ns-policy-based-routing -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-policy-based-routing -- ping -c 4 172.16.1.101`)
	r.Run(`result=$(kubectl exec pods/nettools -n ns-policy-based-routing -- ip r get 172.16.3.1 from 172.16.2.201 ipproto tcp dport 6666)` + "\n" + `echo ${result}` + "\n" + `echo ${result} | grep -E -q "172.16.3.1 from 172.16.2.201 via 172.16.2.200 dev nsm-1 table 1"`)
	r.Run(`result=$(kubectl exec pods/nettools -n ns-policy-based-routing -- ip r get 172.16.3.1 from 172.16.2.201 ipproto tcp sport 5555)` + "\n" + `echo ${result}` + "\n" + `echo ${result} | grep -E -q "172.16.3.1 from 172.16.2.201 dev nsm-1 table 2"`)
	r.Run(`result=$(kubectl exec pods/nettools -n ns-policy-based-routing -- ip r get 172.16.4.1 ipproto udp dport 6666)` + "\n" + `echo ${result}` + "\n" + `echo ${result} | grep -E -q "172.16.4.1 dev nsm-1 table 3 src 172.16.1.101"`)
	r.Run(`result=$(kubectl exec pods/nettools -n ns-policy-based-routing -- ip r get 172.16.4.1 ipproto udp dport 6668)` + "\n" + `echo ${result}` + "\n" + `echo ${result} | grep -E -q "172.16.4.1 dev nsm-1 table 4 src 172.16.1.101"`)
	r.Run(`result=$(kubectl exec pods/nettools -n ns-policy-based-routing -- ip -6 route get 2004::5 from 2004::3 ipproto udp dport 5555)` + "\n" + `echo ${result}` + "\n" + `echo ${result} | grep -E -q "via 2004::6 dev nsm-1 table 5 src 2004::3"`)
}
func (s *Suite) TestScale_from_zero() {
	r := s.Runner("../deployments-k8s/examples/features/scale-from-zero")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-scale-from-zero`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/scale-from-zero?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait -n ns-scale-from-zero --for=condition=ready --timeout=1m pod -l app=nse-supplier-k8s`)
	r.Run(`kubectl wait -n ns-scale-from-zero --for=condition=ready --timeout=1m pod -l app=alpine`)
	r.Run(`kubectl wait -n ns-scale-from-zero --for=condition=ready --timeout=1m pod -l app=nse-icmp-responder`)
	r.Run(`NSE=$(kubectl get pod -n ns-scale-from-zero --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}' -l app=nse-icmp-responder)`)
	r.Run(`kubectl exec pods/alpine -n ns-scale-from-zero -- ping -c 4 169.254.0.0`)
	r.Run(`kubectl exec $NSE -n ns-scale-from-zero -- ping -c 4 169.254.0.1`)
	r.Run(`NSE_NODE=$(kubectl get pod -n ns-scale-from-zero --template '{{range .items}}{{.spec.nodeName}}{{"\n"}}{{end}}' -l app=nse-icmp-responder)` + "\n" + `NSC_NODE=$(kubectl get pod -n ns-scale-from-zero --template '{{range .items}}{{.spec.nodeName}}{{"\n"}}{{end}}' -l app=alpine)`)
	r.Run(`if [ $NSC_NODE == $NSE_NODE ]; then echo "OK"; else echo "different nodes"; false; fi`)
	r.Run(`kubectl delete pod -n ns-scale-from-zero alpine`)
	r.Run(`kubectl wait -n ns-scale-from-zero --for=delete --timeout=1m pod -l app=nse-icmp-responder`)
}
func (s *Suite) TestScaled_registry() {
	r := s.Runner("../deployments-k8s/examples/features/scaled-registry")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-scaled-registry`)
		r.Run(`kubectl scale --replicas=1 deployments/registry-k8s -n nsm-system`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/scaled-registry?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-scaled-registry`)
	r.Run(`NSE=$(kubectl get pod -n ns-scaled-registry --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}' -l app=nse-kernel)`)
	r.Run(`kubectl get nses -A | grep $NSE`)
	r.Run(`kubectl scale --replicas=0 deployments/registry-k8s -n nsm-system`)
	r.Run(`kubectl wait --for=delete --timeout=1m pod -l app=registry -n nsm-system`)
	r.Run(`kubectl get nses -A | grep $NSE`)
	r.Run(`kubectl scale --replicas=2 deployments/registry-k8s -n nsm-system`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=registry -n nsm-system`)
	r.Run(`kubectl scale --replicas=0 deployments/nse-kernel -n ns-scaled-registry`)
	r.Run(`kubectl get nses -A | grep $NSE` + "\n" + `if [[ "$?" == "1" ]]; then echo OK; else echo "nse entry still exists"; false; fi`)
}
func (s *Suite) TestSelect_forwarder() {
	r := s.Runner("../deployments-k8s/examples/features/select-forwarder")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-select-forwarder`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/select-forwarder?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-select-forwarder`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-select-forwarder`)
	r.Run(`kubectl exec pods/alpine -n ns-select-forwarder -- ping -c 4 169.254.0.0`)
	r.Run(`kubectl exec deployments/nse-kernel -n ns-select-forwarder -- ping -c 4 169.254.0.1`)
	r.Run(`kubectl logs pods/alpine -c cmd-nsc -n ns-select-forwarder | grep "my-forwarder-vpp"`)
}
func (s *Suite) TestWebhook() {
	r := s.Runner("../deployments-k8s/examples/features/webhook")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-webhook`)
	})
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/webhook?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=5m pod -l app=nse-kernel -n ns-webhook`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nettools -n ns-webhook`)
	r.Run(`kubectl exec pods/nettools -n ns-webhook -- curl 172.16.1.100:80 | grep -o "<title>Welcome to nginx!</title>"`)
}
func (s *Suite) TestWebhook_smartvf() {
	r := s.Runner("../deployments-k8s/examples/features/webhook-smartvf")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-webhook-smartvf`)
	})
	r.Run(`WH=$(kubectl get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl wait --for=condition=ready --timeout=1m pod ${WH} -n nsm-system`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/features/webhook-smartvf?ref=c1632838c22fdf074abe49ac7f5177eb5937c02b`)
	r.Run(`kubectl wait --for=condition=ready --timeout=5m pod -l app=nse-kernel -n ns-webhook-smartvf`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod postgres-cl -n ns-webhook-smartvf`)
	r.Run(`kubectl exec pods/postgres-cl -n ns-webhook-smartvf -c postgres-cl -- sh -c 'PGPASSWORD=admin psql -h 172.16.1.100 -p 5432 -U admin test'`)
}
