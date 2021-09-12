// Code generated by gotestmd DO NOT EDIT.
package spire

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
)

type Suite struct {
	base.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/interdomain/spire")
	s.T().Cleanup(func() {
		r.Run(`export KUBECONFIG=$KUBECONFIG1 && kubectl delete ns spire` + "\n" + `export KUBECONFIG=$KUBECONFIG2 && kubectl delete ns spire` + "\n" + `export KUBECONFIG=$KUBECONFIG3 && kubectl delete ns spire`)
	})
	r.Run(`[[ ! -z $KUBECONFIG1 ]]`)
	r.Run(`[[ ! -z $KUBECONFIG2 ]]`)
	r.Run(`[[ ! -z $KUBECONFIG3 ]]`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl apply -k cluster1/`)
	r.Run(`kubectl wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-agent`)
	r.Run(`kubectl wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-server`)
	r.Run(`kubectl exec -n spire spire-server-0 -- \` + "\n" + `/opt/spire/bin/spire-server entry create \` + "\n" + `-spiffeID spiffe://nsm.cluster1/ns/spire/sa/spire-agent \` + "\n" + `-selector k8s_sat:cluster:nsm.cluster1 \` + "\n" + `-selector k8s_sat:agent_ns:spire \` + "\n" + `-selector k8s_sat:agent_sa:spire-agent \` + "\n" + `-node`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl apply -k cluster2/`)
	r.Run(`kubectl wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-agent`)
	r.Run(`kubectl wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-server`)
	r.Run(`kubectl exec -n spire spire-server-0 -- \` + "\n" + `/opt/spire/bin/spire-server entry create \` + "\n" + `-spiffeID spiffe://nsm.cluster2/ns/spire/sa/spire-agent \` + "\n" + `-selector k8s_sat:cluster:nsm.cluster1 \` + "\n" + `-selector k8s_sat:agent_ns:spire \` + "\n" + `-selector k8s_sat:agent_sa:spire-agent \` + "\n" + `-node`)
	r.Run(`export KUBECONFIG=$KUBECONFIG3`)
	r.Run(`kubectl apply -k cluster3/`)
	r.Run(`kubectl wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-agent`)
	r.Run(`kubectl wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-server`)
	r.Run(`kubectl exec -n spire spire-server-0 -- \` + "\n" + `/opt/spire/bin/spire-server entry create \` + "\n" + `-spiffeID spiffe://nsm.cluster3/ns/spire/sa/spire-agent \` + "\n" + `-selector k8s_sat:cluster:nsm.cluster1 \` + "\n" + `-selector k8s_sat:agent_ns:spire \` + "\n" + `-selector k8s_sat:agent_sa:spire-agent \` + "\n" + `-node`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1 && bundle1=$(kubectl exec spire-server-0 -n spire -- bin/spire-server bundle show -format spiffe)` + "\n" + `export KUBECONFIG=$KUBECONFIG2 && bundle2=$(kubectl exec spire-server-0 -n spire -- bin/spire-server bundle show -format spiffe)` + "\n" + `export KUBECONFIG=$KUBECONFIG3 && bundle3=$(kubectl exec spire-server-0 -n spire -- bin/spire-server bundle show -format spiffe)`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`echo $bundle2 | kubectl exec -i spire-server-0 -n spire -- bin/spire-server bundle set -format spiffe -id "spiffe://nsm.cluster2"` + "\n" + `echo $bundle3 | kubectl exec -i spire-server-0 -n spire -- bin/spire-server bundle set -format spiffe -id "spiffe://nsm.cluster3"`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`echo $bundle1 | kubectl exec -i spire-server-0 -n spire -- bin/spire-server bundle set -format spiffe -id "spiffe://nsm.cluster1"` + "\n" + `echo $bundle3 | kubectl exec -i spire-server-0 -n spire -- bin/spire-server bundle set -format spiffe -id "spiffe://nsm.cluster3"`)
	r.Run(`export KUBECONFIG=$KUBECONFIG3`)
	r.Run(`echo $bundle1 | kubectl exec -i spire-server-0 -n spire -- bin/spire-server bundle set -format spiffe -id "spiffe://nsm.cluster1"` + "\n" + `echo $bundle2 | kubectl exec -i spire-server-0 -n spire -- bin/spire-server bundle set -format spiffe -id "spiffe://nsm.cluster2"`)
}
func (s *Suite) Test() {}
