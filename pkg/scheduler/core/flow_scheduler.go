/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package core

import (

	"k8s.io/api/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/kubernetes/pkg/scheduler/algorithm"
	schedulercache "k8s.io/kubernetes/pkg/scheduler/cache"
	"k8s.io/kubernetes/pkg/scheduler/core/equivalence"
	schedulerinternalcache "k8s.io/kubernetes/pkg/scheduler/internal/cache"
	internalqueue "k8s.io/kubernetes/pkg/scheduler/internal/queue"
	"k8s.io/kubernetes/pkg/scheduler/volumebinder"
)


// Error returns detailed information of why the pod failed to fit on each node

type flowScheduler struct {
	cache                    schedulerinternalcache.Cache
	equivalenceCache         *equivalence.Cache
	schedulingQueue          internalqueue.SchedulingQueue
	predicates               map[string]algorithm.FitPredicate
	priorityMetaProducer     algorithm.PriorityMetadataProducer
	predicateMetaProducer    algorithm.PredicateMetadataProducer
	prioritizers             []algorithm.PriorityConfig
	extenders                []algorithm.SchedulerExtender
	lastNodeIndex            uint64
	alwaysCheckAllPredicates bool
	cachedNodeInfoMap        map[string]*schedulercache.NodeInfo
	volumeBinder             *volumebinder.VolumeBinder
	pvcLister                corelisters.PersistentVolumeClaimLister
	pdbLister                algorithm.PDBLister
	disablePreemption        bool
	percentageOfNodesToScore int32
}

func (fs *flowScheduler)Schedule(pod *v1.Pod, nl algorithm.NodeLister) (selectedMachine string, err error){
	return "",nil
}

func (fs *flowScheduler)Preempt(*v1.Pod, algorithm.NodeLister, error) (selectedNode *v1.Node, preemptedPods []*v1.Pod, cleanupNominatedPods []*v1.Pod, err error){
	return nil,nil,nil,nil
}

func (fs *flowScheduler)Predicates() map[string]algorithm.FitPredicate{
	return nil
}
// Prioritizers returns a slice of priority config. This is exposed for
// testing.
func (fs *flowScheduler)Prioritizers() []algorithm.PriorityConfig{
	return nil
}